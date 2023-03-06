package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/Banana-Boat/gRPC-template/main-service/internal/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config: ")
	}

	conn, err := sql.Open(
		config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to db: ")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
