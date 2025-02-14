package filemanager

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/13excite/c24-expences/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestFindFiles(t *testing.T) {
	tempDir := t.TempDir()

	testFiles := []string{
		filepath.Join(tempDir, "file1.csv"),
		filepath.Join(tempDir, "file2.csv"),
		filepath.Join(tempDir, "file3.csv"),
	}

	for _, file := range testFiles {
		err := os.MkdirAll(filepath.Dir(file), 0755)
		assert.NoError(t, err)
		_, err = os.Create(file)
		assert.NoError(t, err)
	}

	db := models.DBModel{} // Mock or initialize your DBModel as needed
	// test correct file path
	fileManager := NewFileManager(tempDir, db)

	err := fileManager.findFiles()
	assert.NoError(t, err)

	assert.ElementsMatch(t, testFiles, fileManager.initFiles)

	// test incorrect file path
	fileManager = NewFileManager("wrong/path", db)
	err = fileManager.findFiles()
	assert.Error(t, err)
}

func TestContainsSHA256(t *testing.T) {
	fileManager := &FileManager{}

	files := []models.SHAFile{
		{Path: "file1.csv", SHA256: "ece16ead1f304c31347d26bd0a691ef7eb3962d198fcd28c97936998f5f99345"},
		{Path: "file2.csv", SHA256: "84272daa8e967b93c3f4d71507b15159a606e1d852310ca5ecd9599909f01fdc"},
		{Path: "file3.csv", SHA256: "51b49385a58537be251b45ddc9af64b4322a3ca73aa72e8a0cd336c6696de933"},
	}

	tests := []struct {
		sha256   string
		expected bool
	}{
		{"ece16ead1f304c31347d26bd0a691ef7eb3962d198fcd28c97936998f5f99345", true},
		{"84272daa8e967b93c3f4d71507b15159a606e1d852310ca5ecd9599909f01fdc", true},
		{"51b49385a58537be251b45ddc9af64b4322a3ca73aa72e8a0cd336c6696de933", true},
		{"9f362dd96864c705ca262d1984e34992b117e398e4230630b11646e5d150b71b", false},
	}

	for _, tc := range tests {
		t.Run(tc.sha256, func(t *testing.T) {
			result := fileManager.containsSHA256(files, tc.sha256)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCalculateSHA256(t *testing.T) {
	// Create a temporary file for testing and defer its removal
	tempFile, err := os.CreateTemp("", "test")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write some content to the file
	content := []byte("Hello, World!")
	_, err = tempFile.Write(content)
	assert.NoError(t, err)
	tempFile.Close()

	// test correct file path
	fileManager := &FileManager{}
	sha256, err := fileManager.calculateSHA256(tempFile.Name())
	assert.NoError(t, err)

	// Expected SHA256 hash of "Hello, World!" shasum -a 256 <file>
	expectedSHA256 := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"
	assert.Equal(t, expectedSHA256, sha256)

	// test incorrect file path
	_, err = fileManager.calculateSHA256("wrong/path")
	assert.Error(t, err)
}
