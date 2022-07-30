package main

import (
	"bufio"
	"dataSanitizer/db"
	"dataSanitizer/utils"
	"fmt"
)

const (
	testFile     = "test_access.log"
	completeFile = "access.log"
)

func main() {
	rawData := make([]string, 5)
	data := make([]string, 5)
	file := utils.ReadFile(fmt.Sprintf("%s", testFile))

	// Close file when main function finishes
	defer file.Close()

	// Read file line by line
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		rawData = utils.StringSliceFromRegexFindAll(scn.Text(), `[^\s]+`, 5)
		// ipAddress | requestType | requestPath | responseStatusCode | accessDate

		// save private
		data[1] = utils.FilterNullString(l[1])

		// save incompleto
		data[2] = utils.FilterNullString(l[2])

		// save data da ultima compra
		data[3] = utils.FilterNullString(l[3])

		// save ticket medio
		data[4] = utils.FilterComma(l[4])

		// save ticket da ultima compra
		data[5] = utils.FilterComma(l[5])

		// save loja mais frequente
		data[6] = utils.FilterAndValidateCNPJ(l[6])
		data[6] = string(db.InsertStoreData(data[6]))

		// save loja da ultima compra
		data[7] = utils.FilterAndValidateCNPJ(l[7])
		data[7] = string(db.InsertStoreData(data[7]))

		db.InsertPersonData(data)

		//time.Sleep(2 * time.Second)
	}
}
