package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
)

// This sensitive information must come from environment variables
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "data"
)

func BulkLogDataInsert(rows [][]interface{}) error {
	conn, err := openDBConnection()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	copyCount, err := conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"log_data"},
		[]string{"ip_address", "request_type", "request_path", "response_status_code", "access_date"},
		pgx.CopyFromRows(rows))

	if err != nil {
		return err
	}

	fmt.Printf("%d rows inserted\n", copyCount)

	return nil
}

func openDBConnection() (*pgx.Conn, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", host, port, user, password, dbname)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
