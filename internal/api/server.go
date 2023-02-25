package api

import (
	"fmt"

	"github.com/Banana-Boat/gRPC-template/internal/db"
	"github.com/Banana-Boat/gRPC-template/internal/pb"
	"github.com/Banana-Boat/gRPC-template/internal/util"
)

type Server struct {
	pb.UnimplementedGPRCTemplateServer // 使得还未被实现的rpc能够被接受
	config                             util.Config
	store                              *db.Store
	tokenMaker                         *util.TokenMaker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := util.NewTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
