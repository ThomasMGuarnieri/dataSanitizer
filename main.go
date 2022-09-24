package main

import (
	"dataSanitizer/db"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const threads int = 4

var (
	wg   = sync.WaitGroup{}
	rows [][]interface{}
)

func extractLogLineData(s chan string, r *regexp.Regexp) {
	// Needed data:
	// ipAddress | requestType | requestPath | responseStatusCode | accessDate
	for l := range s {
		match := r.FindStringSubmatch(l)
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

			rows = append(rows, []interface{}{match[1], match[3], match[4], sc, dt})
		}
	}
	wg.Done()
}

func main() {

	f, err := os.ReadFile(fmt.Sprintf("%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	// Changes on the string below may break this program
	rg := regexp.MustCompile(`(.+) - - \[(.+)] "(.+?) /(.+)" (\d{3})`)

	txt := string(f)

	inputCh := make(chan string, 300)
	for i := 0; i < threads; i++ {
		go extractLogLineData(inputCh, rg)
	}

	wg.Add(threads)

	for _, l := range strings.Split(txt, "\n") {
		inputCh <- l
	}
	close(inputCh)
	wg.Wait()
	fmt.Println("Finished processing file")

	err = db.BulkLogDataInsert(rows)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished saving data to database")
}
