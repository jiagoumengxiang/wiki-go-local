---
phase: quick
plan: 260405-dyl
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/handlers/page.go
  - internal/utils/navigation.go
  - internal/handlers/sitemap.go
autonomous: true
requirements: []
user_setup: []

must_haves:
  truths:
    - "Navigation tree displays all .md files as clickable leaf nodes"
    - "Users can access any .md file via URL path (e.g., /folder/file.md)"
    - "Directory listings show both subdirectories and .md files"
    - "Sitemap includes all .md files, not just document.md"
    - "document.md remains the default when accessing a directory"
  artifacts:
    - path: internal/handlers/page.go
      provides: "Route handling for both directory and file paths"
      contains: "Modified PageHandler to handle .md file URLs"
    - path: internal/utils/navigation.go
      provides: "Navigation tree with file-level nodes"
      contains: "BuildNavigation updated to include all .md files"
    - path: internal/handlers/sitemap.go
      provides: "Sitemap with all .md files"
      contains: "SitemapHandler updated to include all .md files"
  key_links:
    - from: "PageHandler"
      to: "os.Stat"
      via: "Check if path is directory or .md file"
      pattern: "os.Stat.*IsDir\\(\\)|strings.HasSuffix.*\\.md"
    - from: "navigation.go BuildNavigation"
      to: "NavItem"
      via: "Add .md files as leaf nodes in tree"
      pattern: "info.IsDir\\(\\).*\\.md"
---

<objective>
Support listing arbitrary MD files and extend file tree to file level

Purpose: Enable flexible document organization by allowing users to access and manage any .md file in the wiki, not just document.md files.

Output: Enhanced navigation tree showing all .md files, updated routing to handle file paths, and sitemap including all documents.
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/todos/completed/2026-04-05-support-listing-all-md-files-extend-file-tree-file-level.md
@internal/handlers/page.go
@internal/utils/navigation.go
@internal/handlers/sitemap.go
@internal/types/types.go
@internal/routes/routes.go

Current limitations:
- PageHandler only handles directory paths, looks for document.md in each directory
- Navigation tree only includes directories, skips all files (lines 96-97 in navigation.go)
- Directory listing only shows directories, skips files (line 175 in page.go)
- Sitemap only includes document.md files (line 211 in sitemap.go)
</context>

<tasks>

<task type="auto">
  <name>Task 1: Update PageHandler to handle file paths and directory listings</name>
  <files>internal/handlers/page.go</files>
  <action>
    Modify PageHandler (lines 103-200) to:
    1. Check if fsPath is a directory OR a .md file (lines 104-108)
    2. If it's a .md file:
       - Read and render the file content directly
       - Parse frontmatter and extract layout
       - Set lastModified from file ModTime
       - Skip directory listing (no dirContent)
    3. If it's a directory:
       - Keep existing document.md logic (lines 128-163)
       - Update directory listing to show both subdirectories AND .md files (lines 165-200):
         * For directories: Show with "is-dir" class
         * For .md files: Show with "is-file" class, link to folder/file.md
         * Extract title from H1 or use formatted filename
         * Skip document.md in file listing (keep as default doc)
         * Use GetDocumentTitle for document.md, FormatFileName for other .md files

    Add helper function FormatFileName (if not exists) to format .md filenames:
    - Remove .md extension
    - Replace underscores/hyphens with spaces
    - Title case the result

    Verify URL path handling works for:
    - /my-folder (directory → shows document.md)
    - /my-folder/my-doc.md (file → shows my-doc.md)
    - /my-folder/other-doc.md (file → shows other-doc.md)
  </action>
  <verify>
    <automated>cd /home/emacs/projects/2.doing/wiki-go-local && go test -run TestPageHandler ./internal/handlers/ 2>/dev/null || echo "No existing tests - manual verification required"</automated>
  </verify>
  <done>
    PageHandler successfully routes to both directory and file paths
    Directory listing shows subdirectories and .md files with appropriate classes
    document.md remains accessible as default for directories
  </done>
</task>

<task type="auto">
  <name>Task 2: Update navigation tree to include all .md files</name>
  <files>internal/utils/navigation.go</files>
  <action>
    Modify BuildNavigation (lines 66-140) to include .md files as leaf nodes:
    1. Update filepath.Walk to process both directories AND .md files (lines 86-94)
       - Change condition to allow .md files: `path == docsPath || (!info.IsDir() && !strings.HasSuffix(path, ".md"))`
    2. Remove the document.md skip logic (lines 96-97) - we now process ALL .md files
    3. For .md files, create NavItem with IsDir=false:
       - Extract title from file's H1 or use formatted filename
       - Set Path to full URL including .md extension (e.g., /folder/file.md)
       - Set IsDir=false
       - No children for files
    4. For directories, keep existing logic (IsDir=true, with children)

    Update GetDocumentTitle (lines 25-48) to also handle .md files directly:
    - Accept filePath parameter (not just dirPath)
    - If filePath ends with .md, read that file directly
    - Extract H1 from the file content
    - Fallback to formatted filename if no H1

    Verify navigation tree includes:
    - Directories with IsDir=true
    - .md files with IsDir=false as leaf nodes
    - document.md shown as a file node (not special handling)
  </action>
  <verify>
    <automated>cd /home/emacs/projects/2.doing/wiki-go-local && go test -run TestBuildNavigation ./internal/utils/ 2>/dev/null || echo "No existing tests - manual verification required"</automated>
  </verify>
  <done>
    Navigation tree displays all .md files as clickable leaf nodes
    Each file node has correct path including .md extension
    H1 titles extracted correctly for file display
  </done>
</task>

<task type="auto">
  <name>Task 3: Update sitemap generation to include all .md files</name>
  <files>internal/handlers/sitemap.go</files>
  <action>
    Modify SitemapHandler (lines 210-260) to include all .md files:
    1. Change filepath.Walk condition to include any .md file (line 211):
       - Old: `!info.IsDir() && filepath.Base(path) == "document.md"`
       - New: `!info.IsDir() && strings.HasSuffix(path, ".md")`
    2. For each .md file:
       - Get the parent directory path for the URL
       - Append the .md filename to the URL path (e.g., /folder/file.md)
       - Check access rules using the full file path
       - Include in sitemap with appropriate metadata

    Extract document title logic (around line 246+) to work with any .md file:
    - Read the file content
    - Extract H1 title or use formatted filename
    - Use title in sitemap if applicable (optional enhancement)

    Verify sitemap XML includes:
    - All .md files with correct URLs
    - Proper last-modified timestamps
    - Access control filtering still applied
  </action>
  <verify>
    <automated>cd /home/emacs/projects/2.doing/wiki-go-local && go test -run TestSitemapHandler ./internal/handlers/ 2>/dev/null || echo "No existing tests - manual verification required"</automated>
  </verify>
  <done>
    Sitemap XML includes all .md files with correct URLs
    Each entry has valid last-modified timestamp
    Access control filtering prevents unauthorized documents in sitemap
  </done>
</task>

</tasks>

<verification>
After all tasks complete:
1. Start the application and verify navigation tree shows .md files
2. Visit /some-folder/file.md and verify it renders correctly
3. Visit /some-folder and verify document.md loads as default
4. Check directory listing shows both folders and .md files
5. Verify /sitemap.xml includes all .md files
</verification>

<success_criteria>
- Navigation tree displays all .md files as clickable links
- Direct access to /folder/file.md works and renders content
- Directory access (/folder) loads document.md as default
- Directory listing shows subdirectories AND .md files
- Sitemap includes all .md files with correct paths
- Backward compatibility maintained (existing directory paths work)
</success_criteria>

<output>
After completion, create `.planning/quick/260405-dyl-md/260405-dyl-SUMMARY.md`
</output>
