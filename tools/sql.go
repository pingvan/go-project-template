package tools

import (
	"fmt"
	"path/filepath"

	embeded "go-project-template/storage"
)

func LoadQuery(filename string) (string, error) {
	data, err := embeded.SQLEmbedFS.ReadFile(filepath.Join("queries", filename))
	if err != nil {
		return "", fmt.Errorf("failed to read SQL file %s: %w", filename, err)
	}

	return string(data), nil
}
