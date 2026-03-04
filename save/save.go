// Package save provides atomic file operations to serialize and deserialize game states or data
// using JSON and Binary (gob) formats, ensuring data safety against interruptions.
package save

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ensureDir ensures that the directory for a given file path exists.
func ensureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// safeWriteToFile writes data atomically to a file to prevent corruption during an unexpected crash.
func safeWriteToFile(filePath string, writeFunc func(f *os.File) error) error {
	if err := ensureDir(filePath); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tempFile := filePath + ".tmp"
	f, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer f.Close()

	if err := writeFunc(f); err != nil {
		os.Remove(tempFile) // clean up on error
		return err
	}

	// Close the file explicitly before renaming to avoid locks on Windows
	if err := f.Close(); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	if err := os.Rename(tempFile, filePath); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename temporary file to intended target: %w", err)
	}

	return nil
}

// SaveJSON serializes the given data of type T and saves it as a JSON file atomically.
// T must be a struct or type that is encodable by the standard json package.
func SaveJSON[T any](filename string, data T) error {
	return safeWriteToFile(filename, func(f *os.File) error {
		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "  ") // format for readability
		if err := encoder.Encode(data); err != nil {
			return fmt.Errorf("json encoding failed: %w", err)
		}
		return nil
	})
}

// LoadJSON reads the given file and deserializes its JSON content back to the type T.
func LoadJSON[T any](filename string) (T, error) {
	var data T
	f, err := os.Open(filename)
	if err != nil {
		return data, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("json decoding failed: %w", err)
	}

	return data, nil
}

// SaveBinary serializes the given data of type T and saves it as a binary file using gob atomically.
// T must be supported by the gob package (e.g. exported fields).
func SaveBinary[T any](filename string, data T) error {
	return safeWriteToFile(filename, func(f *os.File) error {
		encoder := gob.NewEncoder(f)
		if err := encoder.Encode(data); err != nil {
			return fmt.Errorf("binary (gob) encoding failed: %w", err)
		}
		return nil
	})
}

// LoadBinary reads the given file and deserializes its binary gob content back to the type T.
func LoadBinary[T any](filename string) (T, error) {
	var data T
	f, err := os.Open(filename)
	if err != nil {
		return data, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("binary (gob) decoding failed: %w", err)
	}

	return data, nil
}
