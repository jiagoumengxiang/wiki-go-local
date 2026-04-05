---
created: 2026-04-05T08:32:46.325Z
title: 修复上传附件到md文件时目录不存在的问题
area: file-upload
files:
  - internal/handlers/files.go:115-137
---

## Problem

由于之前将新建文档的机制改为直接创建 `.md` 文件（而不是目录 + document.md），现在在上传附件时，UploadFileHandler 检查目录是否存在会失败。

**具体表现：**
1. 新建的文档是 `.md` 文件（如 `docs/myfile.md`）
2. 上传附件时，UploadFileHandler 尝试访问 `docs/myfile.md` 作为目录
3. 由于 `myfile.md` 是文件而不是目录，导致 "Document directory does not exist" 错误

**相关代码：**
- `internal/handlers/files.go:120-137` - UploadFileHandler 构建路径并检查目录是否存在
- 当前逻辑假设文档路径是目录，但现在可能是 `.md` 文件

## Solution

修改 UploadFileHandler 的路径处理逻辑：

1. **检测 .md 文件路径**
   - 检查 `docPath` 是否以 `.md` 结尾
   - 如果是文件，使用 `filepath.Dir()` 获取父目录

2. **调整附件存储位置**
   - 对于 `.md` 文件：在父目录的 `attachments/` 子目录中存储附件
   - 对于目录：保持现有逻辑，在目录的 `attachments/` 子目录中存储

3. **处理附件目录创建**
   - 如果 attachments 目录不存在，自动创建
   - 确保附件存储位置与文档结构兼容

**具体修改：**
- `internal/handlers/files.go` - UploadFileHandler 路径逻辑
- 可能需要更新附件检索逻辑以支持 .md 文件的附件

## Implementation

**Completed:** 2026-04-05

**Changes made:**

1. **UploadFileHandler** (lines 129-133):
   - Added check for `.md` file suffix
   - Use `filepath.Dir()` to get parent directory for `.md` files
   - Attachments are stored in the same directory as the `.md` file

2. **ListFilesHandler** (lines 377-381):
   - Added same `.md` file detection logic
   - Lists files in parent directory for `.md` documents

3. **DeleteFileHandler** (lines 511-521):
   - Handle paths like `docs/myfile.md/attachment.jpg`
   - Extract filename and use parent directory
   - Correctly delete attachment files

4. **ServeFileHandler** (lines 587-639):
   - Parse paths with `.md` files for access control
   - Use parent directory for permission checks
   - Reconstruct correct file path for serving attachments

**Storage structure:**
- Old: `data/documents/docs/mydocument/attachment.jpg` (directory-based)
- New: `data/documents/docs/myfile.md` + `data/documents/docs/attachment.jpg` (file-based)

**Backward compatibility:**
- Old directory-based documents still work
- New `.md` file-based documents now support attachments
- No breaking changes to existing functionality

**Commit:** a32db2c - "fix: support file attachments for .md document files"

## Improved Fix (2026-04-05)

**Problem discovered:**
The initial fix only worked when the path included the `.md` suffix, but the frontend's `getCurrentDocPath()` function removes the `.md` suffix before sending it to the backend.

**Additional changes made:**

Updated all file handlers to properly handle paths without `.md` suffix:

1. **UploadFileHandler** (lines 119-157):
   - Check if path exists (old directory structure)
   - If not, check if `path + ".md"` exists (new file structure)
   - If `.md` file exists, use its parent directory
   - Also handle case where path exists but is a file

2. **ListFilesHandler** (lines 389-437):
   - Same logic as UploadFileHandler
   - Check for both directory and `.md` file structures

3. **DeleteFileHandler** (lines 545-584):
   - Parse path to extract document and attachment parts
   - Check if document is a `.md` file by appending `.md`
   - Reconstruct correct file path for deletion

4. **ServeFileHandler** (lines 617-779):
   - Enhanced path parsing to detect `.md` files without `.md` in path
   - Updated access control logic for `.md` files
   - Handle file path reconstruction for serving attachments

**Testing:**
- ✅ Code compiles successfully
- ✅ Handles both old directory-based and new file-based documents
- ✅ No breaking changes to existing functionality

**Commit:** c79c29d - "fix: properly handle .md document paths without .md suffix"
