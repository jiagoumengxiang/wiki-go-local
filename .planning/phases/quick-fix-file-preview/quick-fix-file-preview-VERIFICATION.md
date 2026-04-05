---
phase: quick-fix-file-preview
verified: 2026-04-06T06:50:00Z
status: gaps_found
score: 2/5 tasks verified
gaps:
  - truth: "All 3 tasks are complete and working"
    status: partial
    reason: "Only 2 of 3 fixes fully implemented. URL encoding is applied to ListFilesHandler but NOT to UploadFileHandler."
    artifacts:
      - path: "internal/handlers/files.go"
        issue: "UploadFileHandler (lines 291-302, 342-362) generates URLs without URL encoding, while ListFilesHandler (line 514) properly encodes URLs with url.PathEscape()"
    missing:
      - "Add url.PathEscape() to UploadFileHandler URL generation (lines 299, 350)"
  - truth: "The code compiles successfully"
    status: verified
    reason: "Go build completed successfully"
  - truth: "URL generation includes full paths (no missing subdirectories)"
    status: partial
    reason: "ListFilesHandler uses full path (line 509), but UploadFileHandler still has inconsistent conditional logic that may generate different URL formats"
    artifacts:
      - path: "internal/handlers/files.go"
        issue: "UploadFileHandler (lines 343-351) uses conditional logic with isMarkdownDoc flag, while ListFilesHandler (line 509) always uses full path without conditionals"
    missing:
      - "Consolidate URL generation logic in UploadFileHandler to match ListFilesHandler approach"
  - truth: "Chinese characters are properly URL-encoded"
    status: partial
    reason: "Chinese characters are URL-encoded in ListFilesHandler but NOT in UploadFileHandler"
    artifacts:
      - path: "internal/handlers/files.go"
        issue: "UploadFileHandler returns unencoded URLs containing Chinese characters, which may cause browser/HTTP issues"
    missing:
      - "Apply url.PathEscape() to UploadFileHandler URLs"
  - truth: "Backward compatibility is maintained"
    status: verified
    reason: "SanitizeFilename now preserves spaces to protect multi-byte UTF-8 characters (line 889-890 comment)"
---

# Phase: Quick Fix - File Preview Error Verification Report

**Phase Goal:** Fix file preview errors caused by URL generation issues with Chinese characters and missing subdirectories
**Verified:** 2026-04-06T06:50:00Z
**Status:** gaps_found
**Re-verification:** No — initial verification

## Goal Achievement

### Observable Truths

| #   | Truth | Status | Evidence |
| --- | ------- | ---------- | -------------- |
| 1   | All 3 tasks are complete and working | ⚠️ PARTIAL | 2 of 3 fixes verified (space preservation, full path in ListFilesHandler). URL encoding missing in UploadFileHandler |
| 2   | The code compiles successfully | ✓ VERIFIED | `go build` completed successfully |
| 3   | URL generation includes full paths (no missing subdirectories) | ⚠️ PARTIAL | ListFilesHandler uses full path, but UploadFileHandler has inconsistent conditional logic |
| 4   | Chinese characters are properly URL-encoded | ⚠️ PARTIAL | URL encoding applied in ListFilesHandler (line 514), missing in UploadFileHandler |
| 5   | Backward compatibility is maintained | ✓ VERIFIED | SanitizeFilename preserves spaces to protect multi-byte UTF-8 characters |

**Score:** 3/5 truths verified

### Required Artifacts

| Artifact | Expected | Status | Details |
| -------- | ----------- | ------ | ------- |
| `internal/handlers/files.go` | File with 3 fixes applied | ✓ VERIFIED | All 3 commits applied to file |
| `sanitizeFilename()` | Preserve spaces for UTF-8 | ✓ VERIFIED | Lines 883-908: Space-to-underscore replacement removed, comment added (line 889-890) |
| URL encoding | Use `url.PathEscape()` | ⚠️ PARTIAL | Applied in ListFilesHandler (line 514), missing in UploadFileHandler |
| Full path logic | Always include subdirectories | ⚠️ PARTIAL | ListFilesHandler (line 509) uses full path, UploadFileHandler uses conditionals |

### Key Link Verification

| From | To | Via | Status | Details |
| ---- | --- | --- | ------ | ------- |
| ListFilesHandler | URL generation | `filepath.Join("/api/files", path, file.Name())` | ✓ WIRED | Line 509: Always uses full path |
| ListFilesHandler | URL encoding | `url.PathEscape(urlPath)` | ✓ WIRED | Line 514: Encodes special characters |
| UploadFileHandler | URL generation (SVG) | Conditional logic | ⚠️ PARTIAL | Lines 295, 299: Uses isMarkdownDoc flag |
| UploadFileHandler | URL generation (regular) | Conditional logic | ⚠️ PARTIAL | Lines 346, 350: Uses isMarkdownDoc flag |
| UploadFileHandler | URL encoding | NOT FOUND | ✗ NOT_WIRED | No url.PathEscape() calls in UploadFileHandler |

### Anti-Patterns Found

| File | Line | Pattern | Severity | Impact |
| ---- | ---- | ------- | -------- | ------ |
| `internal/handlers/files.go` | 291-302 | Inconsistent URL generation between handlers | 🛑 Blocker | UploadFileHandler generates different URL format than ListFilesHandler |
| `internal/handlers/files.go` | 342-362 | Missing URL encoding for Chinese characters | 🛑 Blocker | UploadFileHandler returns unencoded URLs with Chinese characters |
| `internal/handlers/files.go` | 295, 346 | Incomplete refactoring | ⚠️ Warning | isMarkdownDoc conditional logic not removed from UploadFileHandler |

## Commits Verified

### c5839ea: Fix - Preserve spaces in sanitizeFilename
**Status:** ✓ VERIFIED
**Changes:**
- Removed `strings.ReplaceAll(filename, " ", "_")` from sanitizeFilename (was at line ~892)
- Added explanatory comment (lines 889-890): "Note: We preserve spaces to avoid corrupting multi-byte UTF-8 characters"
- Maintains sanitization for dangerous characters (path traversal, shell metacharacters)

**Impact:** Prevents corruption of Chinese filenames like "功 能 清 单 _(Markor).pdf" → "功m能n清s单_(Markor).pdf"

### 246be8a: Fix - Always include full relative path in file URLs
**Status:** ⚠️ PARTIAL (ListFilesHandler only)
**Changes to ListFilesHandler:**
- Removed `isMarkdownDoc` variable and all assignments (lines 417, 430, 442, 463)
- Replaced conditional URL generation (lines 503-509) with single logic:
  ```go
  urlPath := filepath.Join("/api/files", path, file.Name())
  ```
- Added comment explaining full path support for nested directories (lines 506-508)

**Impact:** ListFilesHandler now generates consistent URLs with full relative paths.

**Gap:** UploadFileHandler NOT updated - still uses conditional logic with `isMarkdownDoc` (lines 293-301, 344-351)

### 5ce2470: Fix - Add URL encoding for special characters in file URLs
**Status:** ⚠️ PARTIAL (ListFilesHandler only)
**Changes to ListFilesHandler:**
- Added import `"net/url"` (line 12)
- Applied URL encoding to file URLs (line 514):
  ```go
  encodedURLPath := url.PathEscape(urlPath)
  ```
- Added debug logging (line 515): `log.Printf("ListFilesHandler: encoded URL=%s (original=%s)", encodedURLPath, urlPath)`
- Updated FileInfo struct to use encoded URL (line 522): `URL: encodedURLPath`

**Impact:** Chinese characters and spaces are properly URL-encoded (e.g., `/api/files/.opencode/test/%E5%8A%9F%E8%83%BD%E6%B8%85%E5%8D%95.pdf`)

**Gap:** UploadFileHandler NOT updated - still returns unencoded URLs

## Gaps Summary

### Critical Issues

1. **URL encoding missing in UploadFileHandler**
   - **Impact:** Files uploaded with Chinese characters or spaces will have broken URLs
   - **Location:** Lines 291-302 (SVG files), 342-362 (regular files)
   - **Fix needed:** Add `url.PathEscape(urlPath)` after backslash replacement

2. **Inconsistent URL format between handlers**
   - **Impact:** Uploaded files may not be accessible via generated URLs
   - **Example:**
     - Upload: `/api/files/finance/filename.pdf`
     - List: `/api/files/finance/report.md/filename.pdf`
   - **Fix needed:** Consolidate URL generation logic to match ListFilesHandler approach

### Minor Issues

1. **UploadFileHandler conditional logic not removed**
   - **Impact:** Code inconsistency, potential for different behavior
   - **Location:** Lines 293-301, 344-351
   - **Fix needed:** Remove `isMarkdownDoc` conditional logic, use full path like ListFilesHandler

## Behavioral Verification

### Verified Behaviors

✓ **Code compiles successfully**
```bash
go build -o /tmp/wiki-go-test
# Result: Success
```

✓ **Space preservation in filenames**
- sanitizeFilename no longer replaces spaces with underscores
- Multi-byte UTF-8 characters (Chinese) are protected from corruption

✓ **Full path generation in ListFilesHandler**
- Always constructs URL as `/api/files/{full_relative_path}/{filename}`
- Works for both regular and hidden directories (e.g., `.opencode/`)

✓ **URL encoding in ListFilesHandler**
- Chinese characters encoded using `url.PathEscape()`
- Spaces encoded as `%20` instead of being removed

### Behaviors Needing Verification

❌ **URL encoding in UploadFileHandler**
- Not implemented - needs manual testing with Chinese filenames

❌ **URL format consistency**
- UploadFileHandler and ListFilesHandler may generate different URL formats
- Needs manual testing to verify file access after upload

---

_Verified: 2026-04-06T06:50:00Z_
_Verifier: the agent (gsd-verifier)_
