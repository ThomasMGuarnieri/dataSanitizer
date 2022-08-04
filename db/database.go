package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// This sensitive information must come from environment variables
const (
	host      = "postgres"
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

func Insert(d LogData) (int, error) {
	db, err := openDbConnection()
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId int

	stmt := `insert into data (ip_address, request_type, request_path, response_status_code, access_date) 
		values ($1, $2, $3, $4, $5) returning id`

	err = db.QueryRowContext(ctx, stmt,
		d.IpAddress, d.RequestType, d.RequestPath, d.ResponseStatusCode, d.AccessDate,
	).Scan(&newId)
	if err != nil {
		return 0, nil
	}

	return newId, nil
}

// openConnection opens a new connection with the db
func openDbConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", host, port, user, password, dbname)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
