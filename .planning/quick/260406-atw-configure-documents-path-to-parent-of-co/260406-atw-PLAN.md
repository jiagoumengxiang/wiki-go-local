---
phase: quick
plan: 260406-atw
type: execute
wave: 1
depends_on: []
files_modified: [internal/config/config.go, internal/config/config_test.go]
autonomous: true
requirements: []
user_setup: []

must_haves:
  truths:
    - "When documents_dir is empty, GetDocumentsDir returns parent of config file's directory"
    - "When documents_dir is set, GetDocumentsDir returns root_dir + documents_dir (unchanged)"
    - "Function handles both relative and absolute config file paths correctly"
    - "Tests verify new behavior"
  artifacts:
    - path: internal/config/config.go
      provides: "GetDocumentsDir function with updated logic"
      exports: ["GetDocumentsDir"]
    - path: internal/config/config_test.go
      provides: "Tests verifying GetDocumentsDir uses parent of config directory"
      contains: "TestGetDocumentsDir"
  key_links:
    - from: "internal/config/config.go"
      to: "ConfigFilePath"
      via: "filepath operations"
      pattern: "filepath\\.Dir\\(ConfigFilePath\\)"
---

<objective>
Update GetDocumentsDir to use parent directory of config file when documents_dir is empty

Purpose: Allow documents to be stored in the parent directory of the config file by default, making it easier to organize wiki content separately from the config structure (e.g., config at MDwiki/config.yaml, docs at parent directory)

Output: Modified GetDocumentsDir function and updated tests
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/STATE.md
@internal/config/config.go
@internal/config/config_test.go
</context>

<tasks>

<task type="auto">
  <name>Task 1: Update GetDocumentsDir to use parent of config directory</name>
  <files>internal/config/config.go</files>
  <action>
    Modify the GetDocumentsDir function (lines 369-382) to use the parent directory of the config file instead of the working directory when documents_dir is empty.

    Current behavior:
    - If documents_dir is empty, returns os.Getwd()
    - Falls back to root_dir if os.Getwd() fails

    New behavior:
    - If documents_dir is empty, return the parent directory of the config file's directory
    - Handle both relative and absolute config file paths:
      1. If ConfigFilePath is relative, resolve it against working directory
      2. Get config file's directory using filepath.Dir(configPath)
      3. Return parent of that directory using filepath.Dir(configDir)
    - Keep fallback to root_dir if path resolution fails
    - If documents_dir is not empty, keep existing behavior (root_dir + documents_dir)

    For example:
    - Config at "MDwiki/config.yaml" → return parent directory of "MDwiki" (the working directory or absolute path above MDwiki)
    - Config at "/home/user/wiki/MDwiki/config.yaml" → return "/home/user/wiki"
  </action>
  <verify>grep -A 15 "func GetDocumentsDir" internal/config/config.go | grep -q "filepath.Dir.*ConfigFilePath"</verify>
  <done>GetDocumentsDir returns parent of config directory when documents_dir is empty</done>
</task>

<task type="auto">
  <name>Task 2: Update tests to verify new behavior</name>
  <files>internal/config/config_test.go</files>
  <action>
    Update TestGetDocumentsDir to verify the new behavior:

    1. Modify the "Empty documents_dir uses working directory" test case:
       - Change description to "Empty documents_dir uses parent of config directory"
       - Update expectedBase to be the parent directory of ConfigFilePath
       - For default "MDwiki/config.yaml", expect the parent directory of "MDwiki"

    2. Keep all other test cases unchanged:
       - Non-empty documents_dir uses root_dir + documents_dir
       - Documents_dir with subdirectory
       - RootDir with path separator

    3. Add a new test case to verify absolute config path handling:
       - Test with absolute config path (e.g., "/tmp/test/wiki/config.yaml")
       - Verify it returns "/tmp/test" (parent of config's directory)

    Update TestGetDocumentsDir_ErrorHandling if needed to match new error handling logic.
  </action>
  <verify>go test -v -run TestGetDocumentsDir ./internal/config/</verify>
  <done>Tests pass and verify parent directory behavior</done>
</task>

</tasks>

<verification>
Run full config package tests:
  go test -v ./internal/config/

Verify behavior manually:
  - With config at "MDwiki/config.yaml", GetDocumentsDir with empty documents_dir should return parent of MDwiki
  - With documents_dir set to "docs", should return "MDwiki/docs"
</verification>

<success_criteria>
- GetDocumentsDir returns parent of config directory when documents_dir is empty
- All existing tests pass with updated expectations
- Function handles both relative and absolute config paths
- Backward compatibility maintained for non-empty documents_dir
</success_criteria>

<output>
After completion, create `.planning/quick/260406-atw-configure-documents-path-to-parent-of-co/260406-atw-SUMMARY.md`
</output>
