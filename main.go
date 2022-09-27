package main

import (
	"bufio"
	"dataSanitizer/db"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

// Simultaneous threads running during execution
const threads = 8

var (
	rows  [][]interface{}
	wg    = sync.WaitGroup{}
	mutex = sync.Mutex{}
	rg    = regexp.MustCompile(`(.+) - - \[(.+)] "(.+?) /(.+)" (\d{3})`)
)

func extractLogLineData(s chan string) {
	// Needed data:
	// ipAddress | requestType | requestPath | responseStatusCode | accessDate
	for l := range s {
		match := rg.FindStringSubmatch(l)
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

			// In this case, locks are used because the 'rows' variable is a global variable
			// which means it is shared between all the threads. To avoid a race condition,
			// a lock is used on the variable, so the other threads can't read or write wrong values
			mutex.Lock()
			rows = append(rows, []interface{}{match[1], match[3], match[4], sc, dt})
			mutex.Unlock()
		}
	}
	// Done decrements the wait group counter by one
	wg.Done()
}

func main() {
	// Just OPEN the file using the first argument as the filename
	file, err := os.Open(fmt.Sprintf("%s", os.Args[1]))
	if err != nil {
		log.Panicln(err)
	}

	// Close file when main function finishes
	defer file.Close()

	// Create a scanner to read the file
	scn := bufio.NewScanner(file)

	// Creates a buffered channel
	inputCh := make(chan string, 300)
	for i := 0; i < threads; i++ {
		go extractLogLineData(inputCh)
	}

	wg.Add(threads)

	// Read file line by line
	for scn.Scan() {
		inputCh <- scn.Text()
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
