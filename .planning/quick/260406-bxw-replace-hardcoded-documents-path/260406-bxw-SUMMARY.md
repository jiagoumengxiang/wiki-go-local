---
phase: quick
plan: 260406-bxw
subsystem: config
tags: [config, path-resolution, refactoring]

# Dependency graph
requires: [260406-atw]
provides:
  - All hardcoded documents paths replaced with GetDocumentsDir function calls
  - Consistent use of GetDocumentsDir across all handlers
  - BuildNavigation updated to handle empty documentsDir
affects: [documents, file-handling, navigation, all-handlers]

# Tech tracking
tech-stack:
  added: []
  patterns: [centralized-path-resolution, function-refactoring]

key-files:
  created: []
  modified:
    - internal/handlers/files.go
    - internal/handlers/comments.go
    - internal/handlers/move.go
    - internal/handlers/search.go
    - internal/handlers/sitemap.go
    - internal/handlers/page.go
    - internal/handlers/home.go
    - internal/handlers/editor.go
    - internal/handlers/import.go
    - internal/handlers/links_api.go
    - internal/handlers/versions.go
    - internal/handlers/error.go
    - internal/utils/navigation.go

key-decisions:
  - "Replace all hardcoded filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir) with config.GetDocumentsDir(cfg)"
  - "Update BuildNavigation to handle empty documentsDir parameter"
  - "Add config package import to handlers that didn't have it"

patterns-established:
  - "Centralized path resolution: All documents directory access goes through GetDocumentsDir"
  - "Consistent API: Single source of truth for documents directory path"
  - "Config-aware navigation: BuildNavigation now works with full paths or root+doc"

requirements-completed: []

# Metrics
duration: 25min
completed: 2026-04-06
---

# Quick Task 260406-bxw Summary

**Replaced all hardcoded documents paths with GetDocumentsDir function, enabling consistent use of empty documents_dir configuration across the application**

## Performance

- **Duration:** 25 min
- **Started:** 2026-04-06T10:30:00Z
- **Completed:** 2026-04-06T10:55:00Z
- **Tasks:** 1
- **Files modified:** 13

## Accomplishments

- Replaced 32 occurrences of `filepath.Join(cfg.Wiki.RootDir, cfg.Wiki.DocumentsDir)` with `config.GetDocumentsDir(cfg)` across all handler files
- Updated `BuildNavigation` utility function to handle empty `documentsDir` parameter
- Added `config` package import to 4 handler files that didn't have it
- All code compiles successfully and tests pass

## Task Commits

1. **Task 1: Replace hardcoded paths in handlers** - `b7a4c3f` (refactor)

## Files Created/Modified

- `internal/handlers/files.go` - Replaced 10 hardcoded paths
- `internal/handlers/comments.go` - Replaced 1 hardcoded path, added config import
- `internal/handlers/move.go` - Replaced 1 hardcoded path
- `internal/handlers/search.go` - Replaced 1 hardcoded path
- `internal/handlers/sitemap.go` - Replaced 1 hardcoded path
- `internal/handlers/page.go` - Replaced 2 hardcoded paths, updated BuildNavigation calls
- `internal/handlers/home.go` - Replaced 1 hardcoded path
- `internal/handlers/editor.go` - Replaced 4 hardcoded paths, added config import
- `internal/handlers/import.go` - Replaced 1 hardcoded path
- `internal/handlers/links_api.go` - Replaced 1 hardcoded path, added config import
- `internal/handlers/versions.go` - Replaced 1 hardcoded path
- `internal/handlers/error.go` - Replaced 1 hardcoded path
- `internal/utils/navigation.go` - Updated BuildNavigation to handle empty documentsDir

## Decisions Made

- Chose to update BuildNavigation function to handle empty `documentsDir` rather than creating a new function
- Kept backward compatibility by using conditional logic: if documentsDir is empty, use rootDir directly
- Added config package imports to handlers that were missing it (comments.go, editor.go, links_api.go)

## Deviations from Plan

None - plan executed as intended. All hardcoded paths replaced successfully.

## Issues Encountered

- Initial compilation errors due to incorrect variable names in files.go (filePath vs uploadDir, path vs docPath)
- Fixed by using correct variable names matching the context of each function
- LSP warnings about undefined variables were temporary and resolved after fixes

## User Setup Required

None - this is a refactoring change with no external service configuration needed.

## Next Phase Readiness

- GetDocumentsDir function now used consistently across the entire codebase
- Users can now set `documents_dir: ""` in config.yaml to use the parent directory of the config file
- No blockers or concerns

---
*Phase: quick*
*Plan: 260406-bxw*
*Completed: 2026-04-06*
