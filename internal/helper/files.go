package helper

import (
	"io"
	"os"
)

// CreateEmptyFile creates an empty file in the given size
func CreateEmptyFile(path string, size int64) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	// Creating empty file of given size
	file.Seek(size-1, io.SeekStart)
	file.Write([]byte{0})
	return file, nil
}
