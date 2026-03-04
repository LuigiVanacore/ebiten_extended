package save

import (
	"os"
	"path/filepath"
	"testing"
)

type TestData struct {
	Level int
	Score int
	Name  string
}

func TestSaveLoadJSON(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test_save.json")
	data := TestData{Level: 5, Score: 1000, Name: "Hero"}

	// Test Save
	err := SaveJSON(tmpFile, data)
	if err != nil {
		t.Fatalf("SaveJSON failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Fatalf("file not created: %s", tmpFile)
	}

	// Test Load
	loadedData, err := LoadJSON[TestData](tmpFile)
	if err != nil {
		t.Fatalf("LoadJSON failed: %v", err)
	}

	if loadedData.Level != data.Level || loadedData.Score != data.Score || loadedData.Name != data.Name {
		t.Errorf("loaded data mismatch: got %+v, want %+v", loadedData, data)
	}

	// Test Load non-existent file
	_, err = LoadJSON[TestData]("non_existent_file.json")
	if err == nil {
		t.Errorf("expected error when loading non-existent JSON file")
	}
}

func TestSaveLoadBinary(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test_save.bin")
	data := TestData{Level: 10, Score: 50000, Name: "Master"}

	// Test Save
	err := SaveBinary(tmpFile, data)
	if err != nil {
		t.Fatalf("SaveBinary failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Fatalf("file not created: %s", tmpFile)
	}

	// Test Load
	loadedData, err := LoadBinary[TestData](tmpFile)
	if err != nil {
		t.Fatalf("LoadBinary failed: %v", err)
	}

	if loadedData.Level != data.Level || loadedData.Score != data.Score || loadedData.Name != data.Name {
		t.Errorf("loaded data mismatch: got %+v, want %+v", loadedData, data)
	}

	// Test Load non-existent file
	_, err = LoadBinary[TestData]("non_existent_file.bin")
	if err == nil {
		t.Errorf("expected error when loading non-existent bin file")
	}
}
