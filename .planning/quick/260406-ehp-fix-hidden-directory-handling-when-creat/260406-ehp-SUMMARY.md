---
phase: quick
plan: 260406-ehp
subsystem: utils
tags: [bug-fix, path-sanitization, hidden-directories]
dependency_graph:
  requires: []
  provides: [hidden-directory-creation]
  affects: [document-creation]
tech_stack:
  added: []
  patterns: []
key_files:
  created: []
  modified:
    - path: internal/utils/utils.go
      reason: Updated SanitizePath regex to preserve dots
decisions: []
metrics:
  duration: PT5M
  completed_date: "2026-04-06T10:15:00Z"
  tasks_completed: 1
  files_modified: 1
  lines_added: 1
  lines_removed: 1
  commits: 1
---

# Phase Quick Plan 260406-ehp: Fix Hidden Directory Handling Summary

Fixed the SanitizePath function to preserve dots in directory names, enabling document creation in hidden directories.

## One-Liner

Updated SanitizePath regex pattern from `[^a-zA-Z0-9_\-/]` to `[^a-zA-Z0-9_\-\./]` to include dot (.) as a safe character, allowing hidden directory paths like `.config/notes` to be correctly preserved during document creation.

## Changes Made

### Modified Files

1. **internal/utils/utils.go**
   - Updated regex pattern in `SanitizePath()` function (line 15)
   - Changed from: `regexp.MustCompile(\`[^a-zA-Z0-9_\-/]\`)`
   - Changed to: `regexp.MustCompile(\`[^a-zA-Z0-9_\-\./]\`)`
   - This adds the dot (.) character to the set of allowed safe characters

## Impact

### What This Fixes

- Hidden directory names starting with `.` (e.g., `.config`, `.private`) now preserve their leading dot
- Documents can be created in paths like `.config/mydoc.md` without the dot being replaced with a dash
- Enables organization of system configuration files in hidden directories

### What Remains Protected

- Directory traversal attacks (`../`) are still blocked by the `filepath.Clean()` and prefix removal logic
- Other unsafe characters (special chars, spaces, etc.) are still replaced with dashes
- Path sanitization continues to enforce safe file paths

### Backward Compatibility

- No breaking changes to existing functionality
- Regular paths without dots continue to work exactly as before
- The change only affects paths that contain dots, which previously would have been corrupted

## Deviations from Plan

None - plan executed exactly as written.

## Testing Considerations

To verify this fix works correctly:

1. Create a document with path `.config/test`
2. Verify the file is created at `{documents_dir}/.config/test.md` (not `-config/test.md`)
3. Verify the document is accessible via the UI
4. Verify that path traversal attacks (`../`) are still blocked
5. Verify that regular paths without dots still work correctly

## Self-Check: PASSED

- [x] Commit exists: bb1aebf
- [x] File modified: internal/utils/utils.go
- [x] Regex pattern updated to include dot character
- [x] No build errors
- [x] No breaking changes to existing functionality
