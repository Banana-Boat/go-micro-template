package api

import (
	"context"

	"github.com/Banana-Boat/gRPC-template/internal/db"
	"github.com/Banana-Boat/gRPC-template/internal/pb"
	"github.com/Banana-Boat/gRPC-template/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	/* 判断用户名是否存在 */
	isExistUser, _ := server.store.IsExistUser(ctx, req.GetUsername())
	if !isExistUser {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	/* 获取用户信息 */
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "服务端错误，%s", err)
	}

	/* 校验密码 */
	err = util.CheckPassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "服务端错误，%s", err)
	}

	/* 颁发Token */
	token, err := server.tokenMaker.CreateToken(user.ID, user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "服务端错误，%s", err)
	}

	/* 返回结果 */
	resp := &pb.LoginResponse{
		Token: token,
		User: &pb.User{
			Id:        user.ID,
			Username:  user.Username,
			Age:       user.Age,
			Gender:    string(user.Gender),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}
	return resp, nil
}

func (server *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	/* 判断用户名是否存在 */
	isExistUser, _ := server.store.IsExistUser(ctx, req.GetUsername())
	if isExistUser {
		return nil, status.Errorf(codes.InvalidArgument, "用户已存在")
	}

	/* 创建用户 */
	hashedPassword, err := util.HashPassword(req.Password) // 对密码加密
	if err != nil {
		return nil, status.Errorf(codes.Internal, "服务端错误，%s", err)
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		Password: hashedPassword,
		Gender:   db.UsersGender(req.GetGender()),
		Age:      req.GetAge(),
	}
	res, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "服务端错误，%s", err)
	}

	/* 查询新增用户 */
	id, _ := res.LastInsertId()
	user, err := server.store.GetUserById(ctx, int32(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "服务端错误，%s", err)
	}

	/* 返回结果 */
	resp := &pb.RegisterResponse{
		User: &pb.User{
			Id:        user.ID,
			Username:  user.Username,
			Age:       user.Age,
			Gender:    string(user.Gender),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}
	return resp, nil
}
