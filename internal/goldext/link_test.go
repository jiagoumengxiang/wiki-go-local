package goldext

import (
	"testing"
)

func TestGetDocDir(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Markdown file in hidden directory",
			input:    ".opencode/test.md",
			expected: ".opencode",
		},
		{
			name:     "Markdown file in nested directory",
			input:    "docs/guide/start.md",
			expected: "docs/guide",
		},
		{
			name:     "Markdown file in root",
			input:    "readme.md",
			expected: ".",
		},
		{
			name:     "Empty path",
			input:    "",
			expected: "",
		},
		{
			name:     "Root path",
			input:    "/",
			expected: "",
		},
		{
			name:     "Path with leading slash",
			input:    "/docs/guide/start.md",
			expected: "docs/guide",
		},
		{
			name:     "Path without .md extension",
			input:    "docs/guide/",
			expected: "docs/guide/",
		},
		{
			name:     "Path without .md extension - no trailing slash",
			input:    "docs/guide",
			expected: "docs/guide",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDocDir(tt.input)
			if result != tt.expected {
				t.Errorf("getDocDir(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestResolveLocalPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		docPath  string
		expected string
	}{
		{
			name:     "Attachment in hidden directory",
			path:     "attachment.pdf",
			docPath:  ".opencode/test.md",
			expected: "/api/files/.opencode/attachment.pdf",
		},
		{
			name:     "Attachment with spaces",
			path:     "my document.pdf",
			docPath:  ".opencode/test.md",
			expected: "/api/files/.opencode/my%20document.pdf",
		},
		{
			name:     "Homepage attachment",
			path:     "image.png",
			docPath:  "",
			expected: "/api/files/pages/home/image.png",
		},
		{
			name:     "Nested directory attachment",
			path:     "diagram.svg",
			docPath:  "docs/guide/start.md",
			expected: "/api/files/docs/guide/diagram.svg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveLocalPath(tt.path, tt.docPath)
			if result != tt.expected {
				t.Errorf("resolveLocalPath(%q, %q) = %q, want %q", tt.path, tt.docPath, result, tt.expected)
			}
		})
	}
}

func TestTransformMP4Path(t *testing.T) {
	tests := []struct {
		name      string
		videoPath string
		docPath   string
		expected  string
	}{
		{
			name:      "Video in hidden directory",
			videoPath: "demo.mp4",
			docPath:   ".opencode/test.md",
			expected:  "/api/files/.opencode/demo.mp4",
		},
		{
			name:      "Video with spaces",
			videoPath: "my video.mp4",
			docPath:   ".opencode/test.md",
			expected:  "/api/files/.opencode/my%20video.mp4",
		},
		{
			name:      "Homepage video",
			videoPath: "intro.mp4",
			docPath:   "",
			expected:  "/api/files/pages/home/intro.mp4",
		},
		{
			name:      "Nested directory video",
			videoPath: "tutorial.mp4",
			docPath:   "docs/guide/start.md",
			expected:  "/api/files/docs/guide/tutorial.mp4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransformMP4Path(tt.videoPath, tt.docPath)
			if result != tt.expected {
				t.Errorf("TransformMP4Path(%q, %q) = %q, want %q", tt.videoPath, tt.docPath, result, tt.expected)
			}
		})
	}
}
