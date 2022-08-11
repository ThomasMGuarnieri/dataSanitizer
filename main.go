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

const bufSize = 1024 * 1024

func main() {
	var rows [][]interface{}
	buf := make([]byte, bufSize)
	file := utils.ReadFile(fmt.Sprintf("%s", os.Args[1]))
	// Changes on the string below may break this program
	rg := regexp.MustCompile(`(.+) - - \[(.+)] "(.+?) /(.+)" (\d{3})`)

	// Close file when main function finishes
	defer file.Close()

	// Read file line by line
	scn := bufio.NewScanner(file)
	scn.Buffer(buf, bufSize)

	for scn.Scan() {
		rows = append(rows, extractLogLineData(scn.Text(), rg))
	}

	fmt.Println("Finished processing file")
	err := db.BulkLogDataInsert(rows)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished saving data to database")
}

func extractLogLineData(s string, r *regexp.Regexp) []interface{} {
	// Needed data:
	// ipAddress | requestType | requestPath | responseStatusCode | accessDate
	match := r.FindStringSubmatch(s)
	if match != nil {
		// Converts responseStatusCode to integer
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

		return []interface{}{match[1], match[3], match[4], sc, dt}
	}

	return nil
}
