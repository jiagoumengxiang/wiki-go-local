---
phase: quick
plan: 260406-ehp
type: execute
wave: 1
depends_on: []
files_modified: [internal/utils/utils.go]
autonomous: true
requirements: []
user_setup: []

must_haves:
  truths:
    - Hidden directories starting with '.' preserve their leading dot when creating documents
    - Documents can be created in paths like '.config/mydoc.md'
    - Path sanitization still prevents directory traversal and unsafe characters
  artifacts:
    - path: internal/utils/utils.go
      provides: SanitizePath function that preserves dots in paths
      min_lines: 40
  key_links:
    - from: internal/handlers/editor.go
      to: internal/utils/utils.go
      via: utils.SanitizePath call in CreateDocumentHandler
      pattern: utils\.SanitizePath
---

<objective>
Fix SanitizePath function to preserve dots in directory names, allowing creation of documents in hidden directories.

Purpose: Hidden directories (starting with '.') should be supported for organizing system config files
Output: Updated SanitizePath regex that includes dot as safe character
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@internal/utils/utils.go
@internal/handlers/editor.go
</context>

<tasks>

<task type="auto">
  <name>Task 1: Update SanitizePath regex to preserve dots</name>
  <files>internal/utils/utils.go</files>
  <action>
    In the SanitizePath function (line 15), update the regex pattern from:
    ```go
    re := regexp.MustCompile(`[^a-zA-Z0-9_\-/]`)
    ```
    to:
    ```go
    re := regexp.MustCompile(`[^a-zA-Z0-9_\-\./]`)
    ```

    This change adds the dot (.) to the set of safe characters, preventing hidden directory names from being corrupted.

    After this fix, paths like `.config/notes` will be sanitized correctly to `.config/notes` instead of `-config/notes`.
  </action>
  <verify>grep -n `regexp.MustCompile(\`\[^a-zA-Z0-9_\\\\-\\\\./\]\`)` internal/utils/utils.go</verify>
  <done>
    Regex pattern updated to include dot character as safe
    Hidden directories preserve leading dot when creating documents
  </done>
</task>

</tasks>

<verification>
Test creating a document in a hidden directory:
1. Start the application
2. Create a new document with path `.config/test`
3. Verify the file is created at `{documents_dir}/.config/test.md` (not `config/test.md`)
4. Verify the document is accessible via the UI
</verification>

<success_criteria>
- SanitizePath regex includes dot character
- Hidden directory paths are preserved (e.g., `.config/notes.md` stays `.config/notes.md`)
- Directory traversal protection still works (`../` is still removed)
- No regression: regular paths without dots still work correctly
</success_criteria>

<output>
After completion, create `.planning/quick/260406-ehp-fix-hidden-directory-handling-when-creat/260406-ehp-SUMMARY.md`
</output>
