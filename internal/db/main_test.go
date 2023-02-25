package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Banana-Boat/gin-template/internal/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
