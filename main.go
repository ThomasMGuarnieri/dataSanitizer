package main

import (
	"bufio"
	"dataSanitizer/database"
	"dataSanitizer/utils"
	"github.com/klassmann/cpfcnpj"
	_ "github.com/klassmann/cpfcnpj"
	"log"
)

func main() {
	l := make([]string, 8)
	data := make([]string, 8)
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
		l = utils.StringSliceFromRegexFindAll(scn.Text(), `[^\s]+`, 8)
		// CPF/PRIVATE/INCOMPLETO/DATA DA ÚLTIMA COMPRA/TICKET MÉDIO/TICKET DA ÚLTIMA COMPRA/LOJA MAIS FREQUÊNTE/LOJA DA ÚLTIMA COMPRA

		// CPF
		cpf := cpfcnpj.NewCPF(l[0])
		// If the cpf is not valid jump to the next line
		if ! cpf.IsValid() {
			continue
		}

		// save cpf
		data[0] = string(cpf)

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
		data[6] = string(database.InsertStoreData(data[6]))

		// save loja da ultima compra
		data[7] = utils.FilterAndValidateCNPJ(l[7])
		data[7] = string(database.InsertStoreData(data[7]))

		database.InsertPersonData(data)

		//time.Sleep(2 * time.Second)
	}
}


