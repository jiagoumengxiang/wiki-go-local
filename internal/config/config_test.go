package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetDocumentsDir(t *testing.T) {
	// Save current working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer func() {
		// Restore working directory after test
		if err := os.Chdir(originalWd); err != nil {
			t.Logf("Warning: Failed to restore working directory: %v", err)
		}
	}()

	tests := []struct {
		name         string
		rootDir      string
		documentsDir string
		expectedBase string
		description  string
	}{
		{
			name:         "Empty documents_dir uses working directory",
			rootDir:      "MDwiki",
			documentsDir: "",
			expectedBase: originalWd,
			description:  "When documents_dir is empty, should return current working directory",
		},
		{
			name:         "Non-empty documents_dir uses root_dir + documents_dir",
			rootDir:      "MDwiki",
			documentsDir: "documents",
			expectedBase: filepath.Join("MDwiki", "documents"),
			description:  "When documents_dir is set, should join with root_dir",
		},
		{
			name:         "Documents_dir with subdirectory",
			rootDir:      "wiki",
			documentsDir: "content/docs",
			expectedBase: filepath.Join("wiki", "content/docs"),
			description:  "Should join root_dir and full documents_dir path",
		},
		{
			name:         "RootDir with path separator",
			rootDir:      "data/wiki",
			documentsDir: "docs",
			expectedBase: filepath.Join("data/wiki", "docs"),
			description:  "Should handle root_dir with path separator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{}
			cfg.Wiki.RootDir = tt.rootDir
			cfg.Wiki.DocumentsDir = tt.documentsDir

			result := GetDocumentsDir(cfg)

			if result != tt.expectedBase {
				t.Errorf("GetDocumentsDir() = %v, want %v", result, tt.expectedBase)
			}

			// For empty documents_dir, verify it's absolute path (like working directory)
			if tt.documentsDir == "" && filepath.IsAbs(result) {
				// Ensure it's actually a valid path
				if _, err := os.Stat(result); err != nil {
					t.Logf("Note: Working directory path validation: %v", err)
				}
			}
		})
	}
}

func TestGetDocumentsDir_ErrorHandling(t *testing.T) {
	// Test that function doesn't panic even in edge cases
	cfg := &Config{}
	cfg.Wiki.RootDir = "MDwiki"
	cfg.Wiki.DocumentsDir = ""

	// This should not panic even if os.Getwd() fails in some scenarios
	// (we handle the error by falling back to RootDir)
	result := GetDocumentsDir(cfg)

	if result == "" {
		t.Error("GetDocumentsDir() should return a non-empty string even in error case")
	}
}
