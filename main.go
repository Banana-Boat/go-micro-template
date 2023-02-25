package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Banana-Boat/gRPC-template/internal/api"
	"github.com/Banana-Boat/gRPC-template/internal/db"
	"github.com/Banana-Boat/gRPC-template/internal/pb"
	"github.com/Banana-Boat/gRPC-template/internal/util"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcServer := grpc.NewServer()
	pb.RegisterGPRCTemplateServer(grpcServer, server)
	reflection.Register(grpcServer) // 使得grpc客户端能够了解哪些rpc调用被服务端支持，以及如何调用

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
