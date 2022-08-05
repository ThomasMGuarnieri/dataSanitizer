package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

// This sensitive information must come from environment variables
const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "password"
	dbname    = "data"
	dbTimeout = time.Second * 60
)

type LogData struct {
	id                 int
	IpAddress          string
	RequestType        string
	RequestPath        string
	ResponseStatusCode int
	AccessDate         time.Time
}

func BulkLogDataInsert(d []LogData) error {
	var valueStrings []string
	var valueArgs []interface{}

	db, err := openDbConnection()
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, l := range d {
		valueStrings = append(valueStrings, "($1, $2, $3, $4, $5)")
		valueArgs = append(valueArgs, l.IpAddress)
		valueArgs = append(valueArgs, l.RequestType)
		valueArgs = append(valueArgs, l.RequestPath)
		valueArgs = append(valueArgs, l.ResponseStatusCode)
		valueArgs = append(valueArgs, l.AccessDate)
	}

	stmt := fmt.Sprintf("INSERT INTO log_data (ip_address, request_type, request_path, response_status_code, access_date) VALUES %s", strings.Join(valueStrings, ","))
	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// openConnection opens a new connection with the db
func openDbConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
