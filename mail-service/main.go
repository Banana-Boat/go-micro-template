package main

import (
	"log"
	"net"

	"github.com/Banana-Boat/gRPC-template/mail-service/internal/api"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/pb"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMailServiceServer(grpcServer, server)
	reflection.Register(grpcServer) // 使得grpc客户端能够了解哪些rpc调用被服务端支持，以及如何调用

	listener, err := net.Listen("tcp", config.MailServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
