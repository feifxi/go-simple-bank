package util

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgresDB(connStr string) *pgxpool.Pool {
	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Successfully connected to database!")
	return connPool
}