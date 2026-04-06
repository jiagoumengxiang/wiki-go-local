package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"wiki-go/internal/goldext"
	"wiki-go/internal/types"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NavItem represents a navigation item (directory)
type NavItem struct {
	Title    string
	Path     string
	IsDir    bool
	Children []*NavItem
	IsActive bool
}

// GetDocumentTitle extracts the first H1 title from a markdown file
// If filePath ends with .md, reads that file directly
// Otherwise, looks for document.md in the directory
func GetDocumentTitle(filePath string) string {
	var docPath string

	// Check if filePath is a .md file or a directory
	if strings.HasSuffix(filePath, ".md") {
		docPath = filePath
	} else {
		docPath = filepath.Join(filePath, "document.md")
	}

	file, err := os.Open(docPath)
	if err != nil {
		// If no document.md or can't read it, use file/directory name
		name := filepath.Base(filePath)
		if strings.HasSuffix(name, ".md") {
			return FormatFileName(name)
		}
		return FormatDirName(name)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "# ") {
			title := strings.TrimPrefix(line, "# ")
			// Process emojis in the title
			title = goldext.EmojiPreprocessor(title, "")
			return title
		}
	}

	// If no H1 found, use file/directory name
	name := filepath.Base(filePath)
	if strings.HasSuffix(name, ".md") {
		return FormatFileName(name)
	}
	return FormatDirName(name)
}

// FormatDirName formats a directory name by replacing dashes with spaces and title casing
func FormatDirName(name string) string {
	// Replace dashes with spaces
	name = strings.ReplaceAll(name, "-", " ")

	// Title case the words using cases package
	titleCaser := cases.Title(language.English)
	return titleCaser.String(name)
}

// FormatFileName formats a .md filename by removing the extension, replacing underscores/hyphens with spaces, and title casing
func FormatFileName(name string) string {
	// Remove .md extension
	name = strings.TrimSuffix(name, ".md")

	// Replace underscores and hyphens with spaces
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")

	// Title case the words using cases package
	titleCaser := cases.Title(language.English)
	return titleCaser.String(name)
}

// ToURLPath converts a filesystem path to a URL path
func ToURLPath(path string) string {
	// Convert spaces to dashes
	return strings.ReplaceAll(path, " ", "-")
}

// BuildNavigation builds the navigation structure from the root directory
func BuildNavigation(rootDir string, documentsDir string) (*types.NavItem, error) {
	root := &types.NavItem{
		Title:    "Wiki-Go",
		Path:     "/",
		IsDir:    true,
		Children: make([]*types.NavItem, 0),
	}

	// Create the documents directory path
	var docsPath string
	if documentsDir == "" {
		docsPath = rootDir
	} else {
		docsPath = filepath.Join(rootDir, documentsDir)
	}

	// Check if documents directory exists
	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		// Create documents directory if it doesn't exist
		if err := os.MkdirAll(docsPath, 0755); err != nil {
			return nil, err
		}
	}

	err := filepath.Walk(docsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the documents directory itself
		if path == docsPath {
			return nil
		}

		// Skip non-directories and non-.md files
		if !info.IsDir() && !strings.HasSuffix(path, ".md") {
			return nil
		}

		// Skip the pages/home directory in navigation
		if path == filepath.Join(rootDir, "pages", "home") || path == filepath.Join(rootDir, "pages") {
			return filepath.SkipDir
		}

		// Create relative path for the URL
		relPath := strings.TrimPrefix(path, docsPath)
		relPath = strings.TrimPrefix(relPath, string(os.PathSeparator))
		relPath = filepath.ToSlash(relPath)

		// Split the path into components
		parts := strings.Split(relPath, "/")
		current := root

		if info.IsDir() {
			// It's a directory - build directory structure
			// Get the title from document.md's H1 or fallback to formatted directory name
			title := GetDocumentTitle(path)

			// Build the directory structure
			for i := 0; i < len(parts); i++ {
				// Create URL path with dashes
				urlPath := "/" + ToURLPath(filepath.ToSlash(filepath.Join(parts[:i+1]...)))

				// Look for existing directory at this level
				var found *types.NavItem
				for _, child := range current.Children {
					if child.Path == urlPath {
						found = child
						break
					}
				}

				if found == nil {
					// Create new directory item
					dirTitle := ""
					if i == len(parts)-1 {
						dirTitle = title // Use document.md title for leaf nodes
					} else {
						dirTitle = FormatDirName(parts[i])
					}

					found = &types.NavItem{
						Title:    dirTitle,
						Path:     urlPath,
						IsDir:    true,
						Children: make([]*types.NavItem, 0),
					}
					current.Children = append(current.Children, found)
				}
				current = found
			}
		} else {
			// It's a .md file - create a file node (leaf node)
			// Extract title from file's H1 or use formatted filename
			title := GetDocumentTitle(path)
			if title == FormatDirName(filepath.Base(path)) {
				// If no H1 found, use FormatFileName
				title = FormatFileName(filepath.Base(path))
			}

			// Create URL path including .md extension
			urlPath := "/" + ToURLPath(relPath)

			// Find the parent directory to add this file to
			if len(parts) > 1 {
				parentPath := "/" + ToURLPath(filepath.ToSlash(filepath.Join(parts[:len(parts)-1]...)))
				parent := FindNavItem(root, parentPath)
				if parent != nil {
					// Add the file as a child of the parent directory
					fileNode := &types.NavItem{
						Title:    title,
						Path:     urlPath,
						IsDir:    false,
						Children: nil, // Files have no children
					}
					parent.Children = append(parent.Children, fileNode)
				}
			} else {
				// File is at the root level, add directly to root
				fileNode := &types.NavItem{
					Title:    title,
					Path:     urlPath,
					IsDir:    false,
					Children: nil,
				}
				root.Children = append(root.Children, fileNode)
			}
		}

		return nil
	})

	return root, err
}

// FindNavItem finds a navigation item by its path
func FindNavItem(root *types.NavItem, path string) *types.NavItem {
	if root == nil {
		return nil
	}

	// Clean up the path
	path = strings.TrimSuffix(path, "/")
	if path == "" {
		path = "/"
	}

	if root.Path == path {
		return root
	}

	for _, child := range root.Children {
		if found := FindNavItem(child, path); found != nil {
			return found
		}
	}

	return nil
}

// MarkActiveNavItem marks the active navigation item and its parents
func MarkActiveNavItem(root *types.NavItem, currentPath string) {
	if root == nil {
		return
	}

	// Clean up the path
	currentPath = strings.TrimSuffix(currentPath, "/")
	if currentPath == "" {
		currentPath = "/"
	}

	// Mark this item if it matches
	if root.Path == currentPath {
		root.IsActive = true
	}

	// Mark this item if any child is active
	for _, child := range root.Children {
		MarkActiveNavItem(child, currentPath)
		if child.IsActive {
			root.IsActive = true
		}
	}
}

// FilterNavigation filters the navigation tree based on a predicate function
func FilterNavigation(node *types.NavItem, allow func(path string) bool) *types.NavItem {
	if node == nil {
		return nil
	}

	// Create a new node to avoid modifying the original tree
	newNode := &types.NavItem{
		Title:          node.Title,
		Path:           node.Path,
		IsDir:          node.IsDir,
		IsActive:       node.IsActive,
		DocumentLayout: node.DocumentLayout,
		Children:       make([]*types.NavItem, 0),
	}

	for _, child := range node.Children {
		// Check if child path is allowed
		if allow(child.Path) {
			// Recursively filter children
			filteredChild := FilterNavigation(child, allow)
			if filteredChild != nil {
				newNode.Children = append(newNode.Children, filteredChild)
			}
		}
	}

	return newNode
}
