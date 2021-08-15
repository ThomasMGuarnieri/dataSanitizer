package main

import (
	"bufio"
	"dataSanitizer/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	file := utils.ReadFile("base_teste.txt")

	// Close file when main function finishes
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Read file line by line
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		fmt.Println(utils.StringSliceFromRegexFindAll(scn.Text(), `[^\s]+`, 8))
		time.Sleep(2 * time.Second)
	}
}


