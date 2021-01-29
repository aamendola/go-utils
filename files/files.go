package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateTempFileName ...
func CreateTempFileName(filename string) string {
	tempDir := os.TempDir()
	return filepath.Join(tempDir, filename)
}

// Copy ...
func Copy(sourcePath string, destinationPath string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}
	// Write data to dst
	err = ioutil.WriteFile(destinationPath, data, 0666)
	if err != nil {
		panic(err)
	}
}

// Exists ...
func Exists(path string) bool {
	var result bool
	_, err := os.Stat(path)
	if err != nil {
		result = false
	} else {
		result = true
	}
	return result
}

// Read reads the file named by filename and returns the contents.
func Read(filename string) []byte {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return contents
}

// Remove a file
func Remove(file string) {
	err := os.Remove(file)
	if err != nil {
		panic(fmt.Sprintf("File %v could not be removed", file))
	}
}
