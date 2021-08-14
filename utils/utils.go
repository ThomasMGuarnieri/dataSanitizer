package utils

import (
	"log"
	"os"
	"regexp"
)

// ReadFile uses the os.Open to read a file by the given file name
func ReadFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// StringSliceFromRegexFindAll Returns a slice of strings generated
// from the regexp FindAllString function
func StringSliceFromRegexFindAll(s string, expr string, n int) []string {
	rx, _ := regexp.Compile(expr)
	return rx.FindAllString(s, n)
}