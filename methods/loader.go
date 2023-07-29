package methods

import (
	"fmt"
	"net/http"
	"os"
)

type ExtractMethodInterface interface {
	Execute(file string) ([]string, error)
}

// ExtractMethod is a base struct for implementing the extraction methods.
type ExtractMethod struct {
	ExpectedMimeType string
}

// IsValidMimeTypeOrError checks if the file's MIME type matches the expected MIME type.
func (e *ExtractMethod) IsValidMimeTypeOrError(file string, expectedMimeType string) bool {
	if expectedMimeType == "*" {
		return true
	}

	fileContent, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return false
	}

	mimeType := http.DetectContentType(fileContent)

	return mimeType != expectedMimeType
}
