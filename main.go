package main

import (
	"bufio"
	"dataSanitizer/db"
	"dataSanitizer/utils"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	var rows [][]interface{}
	f := utils.ReadFile(fmt.Sprintf("%s", os.Args[1]))
	// Changes on the string below may break this program
	rg := regexp.MustCompile(`(.+) - - \[(.+)] "(.+?) /(.+)" (\d{3})`)

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		var i []interface{}
		s, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		i = extractLogLineData(s, rg)

		if i != nil {
			rows = append(rows, i)
		}
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
