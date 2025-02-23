package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func ConvertToFile(file multipart.File, fullPath string) (*os.File, error) {
	tempFile, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}

	if _, err = io.Copy(tempFile, file); err != nil {
		tempFile.Close()
		return nil, fmt.Errorf("failed to copy data to file: %v", err)
	}

	if err = tempFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close the file: %v", err)
	}

	finalFile, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to reopen the file: %v", err)
	}

	return finalFile, nil
}

func DeleteLocalFile(filepath string) error {
	return nil
}
