---
mode: quick-full
description: Fix file preview error - uploaded file in hidden directory (.opencode/test) shows "file not found"
---

# Quick Task PLAN: Fix File Preview Error

## Issue Description

When files are uploaded to a hidden directory (e.g., `.opencode/test/`), the generated file URLs are incorrect:
- Missing subdirectory in URL path (e.g., `/api/files/.opencode/file.pdf` instead of `/api/files/.opencode/test/file.pdf`)
- Chinese characters in filenames are corrupted (e.g., `功m能n清s单` instead of `功 能 清 单`)

**Root Cause:**
1. `isMarkdownDoc` flag is incorrectly set for nested directories under hidden folders
2. `sanitizeFilename()` removes spaces from filenames, corrupting multi-byte UTF-8 characters

## User Decisions

### Chinese Filename Encoding
- Use proper URL encoding for Chinese characters and spaces
- Do not strip spaces from filenames (current behavior corrupts multi-byte UTF-8)

### URL Generation Strategy
- Include full relative path (directory + filename) in URL
- Fix missing subdirectory issue (e.g., `.opencode/test/` should be preserved in URL)

### the agent's Discretion
- Use Go's `url.QueryEscape()` or `url.PathEscape()` for encoding
- Preserve original filename characters, only sanitize dangerous characters

## Tasks

### Task 1: Fix filename sanitization to preserve spaces and Chinese characters

**Files:**
- `internal/handlers/files.go` (sanitizeFilename function, lines 890-913)

**Action:**
Modify `sanitizeFilename()` to preserve spaces and non-ASCII characters:
- Remove the line that replaces spaces with underscores: `filename = strings.ReplaceAll(filename, " ", "_")`
- Keep other sanitization for dangerous characters (path traversal, shell metacharacters, etc.)
- Add comment explaining why we preserve spaces (multi-byte UTF-8 character safety)

**Verify:**
```bash
go build ./...
```

**Done:**
- `sanitizeFilename()` preserves spaces and Chinese characters
- Only dangerous characters are removed (`..`, `|`, `;`, etc.)

---

### Task 2: Fix URL generation in ListFilesHandler for nested directories

**Files:**
- `internal/handlers/files.go` (ListFilesHandler function, lines 509-521)

**Action:**
Fix the URL generation logic to always include the full relative path:
1. Remove the conditional logic for `isMarkdownDoc` in URL generation (lines 511-519)
2. Always construct URL as `filepath.Join("/api/files", path, file.Name())` to preserve full path
3. The `path` variable already contains the correct relative path (e.g., `.opencode/test`)
4. Remove unnecessary `filepath.Dir()` calls that were stripping directory information

**Verify:**
```bash
go build ./...
```

**Done:**
- URL always includes full relative path (e.g., `/api/files/.opencode/test/功 能 清 单 _(Markor).pdf`)
- Works for both regular directories and hidden directories (`.opencode/`, etc.)
- Log message shows correct URL with full path

---

### Task 3: Add URL encoding for special characters in URLs

**Files:**
- `internal/handlers/files.go` (ListFilesHandler, around line 521)

**Action:**
Import `net/url` package and apply URL path encoding to generated file URLs:
1. Add import: `"net/url"`
2. After constructing `urlPath` (line 521), apply proper URL encoding:
   - Use `url.PathEscape()` to encode special characters while preserving path structure
   - Example: `encodedURLPath := url.PathEscape(urlPath)`
3. Replace `urlPath` with `encodedURLPath` in the FileInfo struct
4. Add log message showing both original and encoded URL for debugging

**Verify:**
```bash
go build ./...
```

**Done:**
- Chinese characters are properly URL-encoded in the returned file list
- Spaces are encoded as `%20` (not removed)
- File preview URLs work correctly in browser
- Log shows encoding applied

---

## Success Criteria

1. ✅ Files in hidden directories (`.opencode/test/`) generate correct URLs with full path
2. ✅ Chinese filenames are properly encoded (no corruption like `功m能n清s单`)
3. ✅ File preview functionality works for all valid file types
4. ✅ Backward compatibility maintained (existing file URLs continue to work)
5. ✅ Code compiles without errors
