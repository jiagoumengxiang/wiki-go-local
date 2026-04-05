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
