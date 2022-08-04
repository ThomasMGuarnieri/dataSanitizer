package utils

import (
	"log"
	"os"
)

// ReadFile uses the os.Open to read a file by the given file name
func ReadFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
