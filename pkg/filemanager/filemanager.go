package filemanager

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/13excite/c24-expences/pkg/models"
)

// FileManager struct that holds the folder path and the files
type FileManager struct {
	folderPath        string
	initFiles         []string
	deduplicatedFiles []models.SHAFile
	DB                models.DBModel
}

// NewFileManager returns a new FileManager struct
func NewFileManager(folderPath string, db models.DBModel) *FileManager {
	return &FileManager{
		folderPath:        folderPath,
		initFiles:         make([]string, 0),
		deduplicatedFiles: make([]models.SHAFile, 0),
		DB:                db,
	}
}

func (f *FileManager) GetFilesToUpload() ([]models.SHAFile, error) {
	if err := f.deduplicateFiles(); err != nil {
		return nil, err
	}
	return f.deduplicatedFiles, nil
}

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
			f.deduplicatedFiles = append(f.deduplicatedFiles, models.SHAFile{Path: file, SHA256: sha256})
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

func (f *FileManager) containsSHA256(files []models.SHAFile, sha256 string) bool {
	for _, file := range files {
		if file.SHA256 == sha256 {
			return true
		}
	}
	return false
}
