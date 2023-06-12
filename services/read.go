package services

import (
	"io"
	"os"
)

func ReadJson(pathFile string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, _ := os.Open(pathFile)
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	return byteValue, nil
}
