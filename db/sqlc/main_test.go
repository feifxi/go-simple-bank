package sqlc

import (
	"os"
	"testing"

	"github.com/feifxi/go-simple-bank/util"
)

var testStore Store

func TestMain(m *testing.M) {
	connStr := "postgresql://root:rootpass@localhost:5432/simple_bank"
	connPool := util.ConnectPostgresDB(connStr)
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
