# Quick Task 260406-bxw: Replace hardcoded documents path with GetDocumentsDir function

## Quick Task Plan

**Mode:** quick
**Quick ID:** 260406-bxw
**Created:** 2026-04-06

## Task Description

Replace all occurrences of hardcoded `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir)` with calls to `GetDocumentsDir(cfg)` function throughout the codebase. This ensures that the updated logic for empty `documents_dir` (returning parent of config directory) is used consistently across the application.

## Implementation

### Task 1: Replace hardcoded paths in handlers

**Files:**
- `internal/handlers/files.go`
- `internal/handlers/page.go`
- `internal/handlers/home.go`
- `internal/handlers/editor.go`
- `internal/handlers/search.go`
- `internal/handlers/sitemap.go`
- `internal/handlers/move.go`
- `internal/handlers/versions.go`
- `internal/handlers/import.go`
- `internal/handlers/links_api.go`
- `internal/handlers/comments.go`
- `internal/handlers/error.go`

**Action:**
Replace `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir)` with `GetDocumentsDir(cfg)` in all occurrences.

**Special cases:**
- For paths that include additional subdirectories (e.g., `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir, path)`), replace with `filepath.Join(GetDocumentsDir(cfg), path)`

**Verify:**
- No occurrences of `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir)` remain in handler files
- Code compiles successfully: `go build ./...`

**Done:**
- All hardcoded paths replaced with GetDocumentsDir function call
- Code compiles without errors

---

## Summary

This is a refactoring task to ensure consistent use of the `GetDocumentsDir` function. The function was recently updated to return the parent directory of the config file when `documents_dir` is empty, but many places in the codebase still use the old hardcoded path pattern.

The task involves a find-and-replace operation across multiple handler files:
1. Find all `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir)` patterns
2. Replace with `GetDocumentsDir(cfg)` for base directory
3. For paths with additional segments, use `filepath.Join(GetDocumentsDir(cfg), ...)`

This will enable users to set `documents_dir: ""` in config to have documents stored in the parent directory of the config file, while keeping the config file in a subdirectory like `MDwiki/`.
