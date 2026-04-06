# Quick Task 260406-cvq: Fix attachment preview path to use config

## Quick Task Plan

**Mode:** quick
**Quick ID:** 260406-cvq
**Created:** 2026-04-06

## Task Description

Fix attachment preview functionality to use the configured documents directory instead of hardcoded paths. The issue was that `goldext/link.go` used hardcoded paths like `"data/documents/"` for attachment resolution, which didn't respect the `documents_dir` configuration.

## Implementation

### Task 1: Add config-aware attachment path resolution

**Files:**
- `internal/goldext/link.go`
- `main.go`

**Action:**
1. Add `GetDocumentsDirFunc` variable to goldext package to store function for getting documents directory
2. Modify `getAttachmentPath` to use this function instead of hardcoded paths
3. Set `GetDocumentsDirFunc` in main.go after config is loaded
4. Handle both regular documents and homepage (pages/home) paths correctly

**Verify:**
- Code compiles successfully
- Attachment preview works with empty documents_dir config
- Attachment preview works with non-empty documents_dir config

**Done:**
- Attachment preview uses configured documents directory
- All path resolution centralized through GetDocumentsDir

---

## Summary

This is a bug fix to make attachment preview work correctly with the new `GetDocumentsDir` configuration. The issue was that the markdown preprocessor in `goldext/link.go` was using hardcoded paths to resolve attachment files during preview.

The fix:
1. Added a callback function `GetDocumentsDirFunc` to goldext package
2. Updated `getAttachmentPath` to use this callback for dynamic path resolution
3. Set the callback in main.go after configuration is loaded

This ensures that attachment previews work correctly whether `documents_dir` is empty (uses parent of config) or set to a specific directory.
