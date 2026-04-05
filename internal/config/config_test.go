package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetDocumentsDir(t *testing.T) {
	// Save current working directory and ConfigFilePath
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	originalConfigPath := ConfigFilePath
	defer func() {
		// Restore working directory after test
		if err := os.Chdir(originalWd); err != nil {
			t.Logf("Warning: Failed to restore working directory: %v", err)
		}
		// Restore ConfigFilePath
		ConfigFilePath = originalConfigPath
	}()

	// Calculate expected parent directory for default config
	// Default ConfigFilePath is "MDwiki/config.yaml", so parent should be parent of "MDwiki"
	defaultConfigDir := filepath.Join(originalWd, "MDwiki")
	expectedParentOfConfigDir := filepath.Dir(defaultConfigDir)

	tests := []struct {
		name         string
		rootDir      string
		documentsDir string
		configPath   string
		expectedBase string
		description  string
	}{
		{
			name:         "Empty documents_dir uses parent of config directory",
			rootDir:      "MDwiki",
			documentsDir: "",
			configPath:   "MDwiki/config.yaml",
			expectedBase: expectedParentOfConfigDir,
			description:  "When documents_dir is empty, should return parent of config's directory",
		},
		{
			name:         "Empty documents_dir with absolute config path",
			rootDir:      "MDwiki",
			documentsDir: "",
			configPath:   "/tmp/test/wiki/config.yaml",
			expectedBase: "/tmp/test",
			description:  "With absolute config path, should return parent of config's directory",
		},
		{
			name:         "Non-empty documents_dir uses root_dir + documents_dir",
			rootDir:      "MDwiki",
			documentsDir: "documents",
			configPath:   "MDwiki/config.yaml",
			expectedBase: filepath.Join("MDwiki", "documents"),
			description:  "When documents_dir is set, should join with root_dir",
		},
		{
			name:         "Documents_dir with subdirectory",
			rootDir:      "wiki",
			documentsDir: "content/docs",
			configPath:   "MDwiki/config.yaml",
			expectedBase: filepath.Join("wiki", "content/docs"),
			description:  "Should join root_dir and full documents_dir path",
		},
		{
			name:         "RootDir with path separator",
			rootDir:      "data/wiki",
			documentsDir: "docs",
			configPath:   "MDwiki/config.yaml",
			expectedBase: filepath.Join("data/wiki", "docs"),
			description:  "Should handle root_dir with path separator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set ConfigFilePath for this test
			ConfigFilePath = tt.configPath

			cfg := &Config{}
			cfg.Wiki.RootDir = tt.rootDir
			cfg.Wiki.DocumentsDir = tt.documentsDir

			result := GetDocumentsDir(cfg)

			if result != tt.expectedBase {
				t.Errorf("GetDocumentsDir() = %v, want %v", result, tt.expectedBase)
			}

			// For empty documents_dir, verify it's absolute path
			if tt.documentsDir == "" && filepath.IsAbs(result) {
				// Ensure it's actually a valid path (unless it's a test path like /tmp/test)
				if _, err := os.Stat(result); err != nil && !os.IsNotExist(err) {
					t.Logf("Note: Path validation: %v", err)
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

	// Verify it returns root_dir as fallback when it can't determine parent directory
	// (this happens when parentDir is "" or "." after path resolution)
	if result != cfg.Wiki.RootDir && result != filepath.Dir(filepath.Dir(filepath.Join(os.TempDir(), ConfigFilePath))) {
		// If it's not RootDir, it should be some valid parent directory
		t.Logf("GetDocumentsDir() returned: %v (RootDir: %v)", result, cfg.Wiki.RootDir)
	}
}
