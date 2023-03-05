package main

import (
	"net"

	"github.com/Banana-Boat/gRPC-template/mail-service/internal/api"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/pb"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/util"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/worker"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	/* 加载配置 */
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config: ")
	}

	/* 创建 Redis 的 Distributor & Processor */
	redisOPt := asynq.RedisClientOpt{
		Addr: config.RedisServerAddress,
	}
	taskDistributor := worker.NewTaskDistributor(redisOPt)
	go runTaskProcessor(redisOPt) // 创建 go routine

	/* 运行 gRPC 服务 */
	runGRPCServer(config, taskDistributor)
}

func runGRPCServer(config util.Config, taskDistributor *worker.TaskDistributor) {
	server, err := api.NewServer(config, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMailServiceServer(grpcServer, server)
	reflection.Register(grpcServer) // 使得grpc客户端能够了解哪些rpc调用被服务端支持，以及如何调用

	listener, err := net.Listen("tcp", config.MailServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt) {
	taskProcessor := worker.NewTaskProcessor(redisOpt)
	log.Info().Msg("start task processor")

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
