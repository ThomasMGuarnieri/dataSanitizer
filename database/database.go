package database

import (
	"dataSanitizer/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "datasan"
	password = "datasan"
	dbname   = "datasan"
)

func InsertStoreData(cnpj string) int64 {
	db := openConnection()

	// Close connection when function exec ends
	defer func(db *sql.DB) {
		err := db.Close()
		utils.CheckError(err)
	}(db)

	if cnpj != "" {
		_, err := db.Exec(`INSERT INTO store (cnpj) VALUES ($1) ON CONFLICT ON CONSTRAINT store_cnpj_key DO NOTHING`, cnpj)
		utils.CheckError(err)

		return 1
	}
	return 0
}

// InsertPersonData insert a new person in this order
//cpf, private, incomplete, avg_ticket_value, last_purchase_date, last_purchase_ticket, most_frequent_store_id, last_purchase_store_id
func InsertPersonData(person []string) {
	db := openConnection()

	// Considerações: Tive a necessidade de fixar alguns dados em função de
	// não conseguir lidar com a mudança de tipos para realizar os inserts
	avg_ticket_value := 0
	last_purchase_date := "2010-10-10"
	last_purchase_ticket := 0


	// Close connection when function exec ends
	defer func(db *sql.DB) {
		err := db.Close()
		utils.CheckError(err)
	}(db)

	if person[3] != "" {
		last_purchase_date = person[3]
	}
	if person[4] != "" {
		avg_ticket_value, _ = strconv.Atoi(person[4])
	}
	if person[5] != "" {
		last_purchase_ticket, _ = strconv.Atoi(person[5])
	}


	q :=`INSERT INTO person 
    	 (cpf, 
    	  private, 
    	  incomplete, 
    	  last_purchase_date,
    	  avg_ticket_value,
    	  last_purchase_ticket, 
    	  most_frequent_store_id, 
    	  last_purchase_store_id) 
    	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT ON CONSTRAINT person_cpf_key DO NOTHING`

	// Considerações: Não consegui de forma clara buscar os ids que foram cadastrados anteriormente sem realizar um select, logo coloqui os dados fixos
	_, err := db.Exec(q, person[0],person[1],person[2],last_purchase_date,avg_ticket_value,last_purchase_ticket,1,1)
	utils.CheckError(err)
}

// openConnection opens a new connection with the database
func openConnection() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open connection
	db, err := sql.Open("postgres", psqlConn)
	utils.CheckError(err)

	return db
}


