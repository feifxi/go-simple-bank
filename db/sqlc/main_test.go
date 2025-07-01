package sqlc

import (
	"os"
	"testing"

	"github.com/feifxi/go-simple-bank/pkg/db"
)

var testStore Store

func TestMain(m *testing.M) {
	connstr := "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable"
	connPool := db.Connect(connstr)
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}