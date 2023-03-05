package main

import (
	"database/sql"
	"log"

	"github.com/Banana-Boat/gRPC-template/main-service/internal/api"
	"github.com/Banana-Boat/gRPC-template/main-service/internal/db"
	"github.com/Banana-Boat/gRPC-template/main-service/internal/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.MainServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
