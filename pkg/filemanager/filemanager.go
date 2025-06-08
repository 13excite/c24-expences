// Package filemanager provides the functionality to find files in a given directory
// and calculate their SHA256 hashes. It then checks if the hash is already in the database
package filemanager

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/13excite/c24-expense/pkg/models"
)

// DBModel is the interface for the database model
type DBModel interface {
	GetSHAFiles() ([]models.SHAFile, error)
	InsertSHAFile(models.SHAFile) error
}

// FileManager struct that holds the folder path and the files
type FileManager struct {
	folderPath        string
	initFiles         []string
	deduplicatedFiles []models.SHAFile
	DB                DBModel
}

// NewFileManager returns a new FileManager struct
func NewFileManager(folderPath string, db DBModel) *FileManager {
	return &FileManager{
		folderPath:        folderPath,
		initFiles:         make([]string, 0),
		deduplicatedFiles: make([]models.SHAFile, 0),
		DB:                db,
	}
}

// GetFilesToUpload returns the files that are not uploaded to database yet
func (f *FileManager) GetFilesToUpload() ([]models.SHAFile, error) {
	if err := f.deduplicateFiles(); err != nil {
		return nil, err
	}
	return f.deduplicatedFiles, nil
}

// deduplicateFiles finds the files in the given directory and calculates
// their SHA256 hashes. It then checks if the hash is already in the database
// and if not, adds the file to the list of files to be uploaded.
func (f *FileManager) deduplicateFiles() error {
	if err := f.findFiles(); err != nil {
		return err
	}

	processedFiles, err := f.DB.GetSHAFiles()
	if err != nil {
		return err
	}
	for _, file := range f.initFiles {
		sha256, err := f.calculateSHA256(file)
		if err != nil {
			return err
		}
		if !f.containsSHA256(processedFiles, sha256) {
			shaFile := models.SHAFile{Path: file, SHA256: sha256}
			f.DB.InsertSHAFile(shaFile)
			f.deduplicatedFiles = append(f.deduplicatedFiles, shaFile)
		}
	}
	return nil
}

// findFiles recursively finds all files in the given directory.
func (f *FileManager) findFiles() error {
	err := filepath.Walk(f.folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip dirs
		if !info.IsDir() {
			f.initFiles = append(f.initFiles, path)
		}
		return nil
	})
	return err
}

// CalculateSHA256 computes the SHA256 hash of a given file.
func (f *FileManager) calculateSHA256(filePath string) (string, error) {
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

// containsSHA256 checks if the given SHA256 hash is in the list of files.
func (f *FileManager) containsSHA256(files []models.SHAFile, sha256 string) bool {
	for _, file := range files {
		if file.SHA256 == sha256 {
			return true
		}
	}
	return false
}
