package main

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/Banana-Boat/gRPC-template/main-service/internal/api"
	"github.com/Banana-Boat/gRPC-template/main-service/internal/db"
	"github.com/Banana-Boat/gRPC-template/main-service/internal/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config: ")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	err = server.Start(config.MainServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server: ")
	}
}
