package main

import (
	"bufio"
	"dataSanitizer/db"
	"dataSanitizer/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	testFile     = "test_access.log"
	completeFile = "access.log"
)

func main() {
	file := utils.ReadFile(fmt.Sprintf("%s", completeFile))
	// Changes on the string below may break this program
	rg := regexp.MustCompile(`(.+) - - \[(.+)] "(.+) /(.+)" (\d{3})`)
	var ld []db.LogData

	// Close file when main function finishes
	defer file.Close()

	// Read file line by line
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		// Needed data:
		// ipAddress | accessDate | requestType | requestPath | responseStatusCode |
		match := rg.FindStringSubmatch(scn.Text())
		if match != nil {
			sc, err := strconv.Atoi(match[5])
			if err != nil {
				fmt.Println("Fail on string to int conversion")
				panic(err)
			}

			dt, err := time.Parse("02/Jan/2006:03:04:05 -0700", match[2])
			if err != nil {
				fmt.Println("Fail on string to date conversion")
				panic(err)
			}

			l := db.LogData{
				IpAddress:          match[1],
				AccessDate:         dt,
				RequestType:        match[3],
				RequestPath:        match[4],
				ResponseStatusCode: sc,
			}

			ld = append(ld, l)
		}
	}

	fmt.Println("Finished processing file")
}
