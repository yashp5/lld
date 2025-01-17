package filestorage

import (
	"testing"
)

func TestFileStorage(t *testing.T) {
	// Create a new FileStorage instance
	fs := NewFileStorage()

	err := fs.Upload("file1.txt", 100)
	if err != nil {
		t.Fatalf("Upload failed: %v", err)
	}

	size, err := fs.Get("file1.txt")
	if err != nil || size != 100 {
		t.Fatalf("Expected size: %d, got: %d (err: %v)", 100, size, err)
	}

	err = fs.Upload("file2.txt", 200)
	if err != nil {
		t.Fatalf("Upload failed: %v", err)
	}

	err = fs.Copy("file1.txt", "file2.txt")
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}
	if fs.hashmap["file1.txt"].Size != fs.hashmap["file2.txt"].Size {
		t.Fatalf("Expected copied file size: %d, got: %d", fs.hashmap["file1.txt"].Size, fs.hashmap["file2.txt"].Size)
	}

	err = fs.Upload("file3.txt", 150)
	if err != nil {
		t.Fatalf("Upload failed: %v", err)
	}

	results, err := fs.Search("file")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	expected := []string{"file3.txt", "file1.txt", "file2.txt"}
	if len(results) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(results))
	}

	for i, name := range expected {
		if results[i] != name {
			t.Errorf("Expected result %d to be %s, got %s", i, name, results[i])
		}
	}
}
