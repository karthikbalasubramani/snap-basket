package grpcdefinedservers

import (
	"context"

	"github.com/karthikbalasubramani/snap-basket/api/pb"
	"github.com/karthikbalasubramani/snap-basket/internal/handlers"
	logx "github.com/karthikbalasubramani/snap-basket/internal/logger"
)

// Loggers
var Info = logx.CustomLogger.Info
var Error = logx.CustomLogger.Error
var Warn = logx.CustomLogger.Warn

type UserServer struct {
	pb.CreateUserModelServiceServer
	pb.LoginUserServiceServer
}

// Create user server function
func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	Info.Println("User Creation Request Received")
	return handlers.HandlerCreateUser(req)
}

// Login Uset server function
func (s *UserServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	Info.Println("Login Request Received")
	return handlers.HandlerLoginUser(req)
}
