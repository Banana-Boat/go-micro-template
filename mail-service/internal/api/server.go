package api

import (
	"context"

	"github.com/Banana-Boat/gRPC-template/mail-service/internal/pb"
	"github.com/Banana-Boat/gRPC-template/mail-service/internal/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedMailServiceServer // 使得还未被实现的rpc能够被接受
	config                            util.Config
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	return server, nil
}

func (Server *Server) SendMail(ctx context.Context, req *pb.SendMailRequest) (*pb.SendMailResponse, error) {
	// 待实现
	// destAddr := req.GetDestAddr()
	resp := &pb.SendMailResponse{
		CreatedAt: timestamppb.Now(),
	}
	return resp, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
