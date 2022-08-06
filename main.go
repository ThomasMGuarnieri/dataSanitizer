package main

import (
	"bufio"
	"dataSanitizer/db"
	"dataSanitizer/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	file := utils.ReadFile(fmt.Sprintf("%s", os.Args[1]))
	// Changes on the string below may break this program
	rg := regexp.MustCompile(`(.+) - - \[(.+)] "(.+?) /(.+)" (\d{3})`)
	var rows [][]interface{}

	// Close file when main function finishes
	defer file.Close()

	// Read file line by line
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		// Needed data:
		// ipAddress | requestType | requestPath | responseStatusCode | accessDate
		match := rg.FindStringSubmatch(scn.Text())
		if match != nil {
			sc, err := strconv.Atoi(match[5])
			if err != nil {
				fmt.Println("Fail on string to int conversion")
				panic(err)
			}

			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[2])
			if err != nil {
				fmt.Println("Fail on string to date conversion")
				panic(err)
			}

			rows = append(rows, []interface{}{match[1], match[3], match[4], sc, dt})
		}
	}

	fmt.Println("Finished processing file")
	err := db.BulkLogDataInsert(rows)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished saving data to database")
}
