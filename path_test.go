package eutil

import (
	"path/filepath"
	"testing"
)

func TestParseFilePath(t *testing.T) {
	tests := []struct {
		input      string
		expectName string
		expectExt  string
	}{
		{"test.txt", "test", ".txt"},
		{"dir/subdir/data.json", "data", ".json"},
		{"/absolute/path/to/main.go", "main", ".go"},
		{"no_extension", "no_extension", ""},
		{".hiddenfile", ".hiddenfile", ""},
		{"./file.with.multiple.dots.tar.gz", "file.with.multiple.dots.tar", ".gz"},
		{"./path with space/file name.txt", "file name", ".txt"},
		{"中文路径/文件.doc", "文件", ".doc"},
		{".", "", ""},
		{"/", "", ""},
	}

	for _, tt := range tests {
		dir, filename, ext, err := ParseFilePath(tt.input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if filename != tt.expectName {
			t.Errorf("filename mismatch: got %q, want %q", filename, tt.expectName)
		}

		if ext != tt.expectExt {
			t.Errorf("extension mismatch: got %q, want %q", ext, tt.expectExt)
		}

		if tt.input != "." && tt.input != "/" && !filepath.IsAbs(dir) {
			t.Errorf("dir should be absolute: got %s", dir)
		}
	}
}
