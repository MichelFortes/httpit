package config

import (
	"os"
	"testing"
)

func TestGetTestScheme_FileNotFound(t *testing.T) {
	_, err := GetTestScheme("non-existent-file.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestGetTestScheme_InvalidJSON(t *testing.T) {
	// Create temp file with bad content
	tmpfile, err := os.CreateTemp("", "bad-json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte("{ invalid json ")); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	_, err = GetTestScheme(tmpfile.Name())
	if err == nil {
		t.Error("Expected error for invalid json, got nil")
	}
}
