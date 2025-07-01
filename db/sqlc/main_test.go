package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store


func TestMain(m *testing.M) {
	dsn := "host=localhost port=5432 dbname=simple_bank user=root password=rootpass sslmode=disable"

	connPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())

}