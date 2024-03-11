package filecrypt

import (
	"io"
	"os"
)

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	text, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return text, nil
}
