package analyze

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAllGoFiles(t *testing.T) {
	// Setup temporary directory structure
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create mock files and directories
	files := []struct {
		path     string
		isDir    bool
		contents string
	}{
		{path: "file1.go", isDir: false, contents: "package main"},
		{path: "file2.go", isDir: false, contents: "package main"},
		{path: "subdir", isDir: true},
		{path: "subdir/file3.go", isDir: false, contents: "package main"},
		{path: "subdir/file4.txt", isDir: false, contents: "Not a Go file"},
	}

	for _, file := range files {
		fullPath := filepath.Join(tempDir, file.path)
		if file.isDir {
			err := os.MkdirAll(fullPath, 0755)
			if err != nil {
				t.Fatalf("Failed to create dir: %v", err)
			}
		} else {
			err := ioutil.WriteFile(fullPath, []byte(file.contents), 0644)
			if err != nil {
				t.Fatalf("Failed to create file: %v", err)
			}
		}
	}

	// Run the function
	goFiles, err := getAllGoFiles(tempDir)
	if err != nil {
		t.Fatalf("Error getting Go files: %v", err)
	}

	// Validate the results
	expectedFiles := []string{
		filepath.Join(tempDir, "file1.go"),
		filepath.Join(tempDir, "file2.go"),
		filepath.Join(tempDir, "subdir/file3.go"),
	}

	if len(goFiles) != len(expectedFiles) {
		t.Fatalf("Expected %d Go files, got %d", len(expectedFiles), len(goFiles))
	}

	for _, expectedFile := range expectedFiles {
		found := false
		for _, goFile := range goFiles {
			if goFile == expectedFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in result", expectedFile)
		}
	}
}
