package filemanager

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// ProcessedFilesSet is a set of processed files for preventing duplicate uploads
type ProcessedFilesSet map[string]struct{}

func (s ProcessedFilesSet) Has(item string) bool {
	_, ok := s[item]
	return ok
}

func (s ProcessedFilesSet) Add(item string) {
	s[item] = struct{}{}
}

// FileManager struct that holds the folder path and the files
type FileManager struct {
	folderPath     string
	files          []string
	ProcessedFiles ProcessedFilesSet
}

// NewFileManager returns a new FileManager struct
func NewFileManager(folderPath string) *FileManager {
	return &FileManager{
		folderPath:     folderPath,
		files:          make([]string, 0),
		ProcessedFiles: make(ProcessedFilesSet),
	}
}

// FindFiles recursively finds all files in the given directory.
func (f *FileManager) FindFiles() error {
	err := filepath.Walk(f.folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip dirs
		if !info.IsDir() {
			f.files = append(f.files, path)
		}
		return nil
	})
	return err
}

// CalculateSHA256 computes the SHA256 hash of a given file.
func CalculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
