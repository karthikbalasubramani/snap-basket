package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/karthikbalasubramani/snap-basket/api/pb"
	logx "github.com/karthikbalasubramani/snap-basket/internal/logger"
	repo "github.com/karthikbalasubramani/snap-basket/internal/repository"
	server "github.com/karthikbalasubramani/snap-basket/servers"
	"google.golang.org/grpc"
)

// Global Declarations for logger
var Info = logx.CustomLogger.Info
var Error = logx.CustomLogger.Error
var Warn = logx.CustomLogger.Warn

// Environmental Values Configuration Go Server
type ServerConfig struct {
	Port     string
	Protocol string
}

// Environmental Values Configuration Database
type DatabaseConfig struct {
	Uri            string
	DatabaseName   string
	UserCollection string
}

// loadenvvariable function to loads environment variables
// If the file doesn't found then printing warning statement
// If the environment variables are not present added default values
func loadenvvariable() ServerConfig {
	err := godotenv.Load()
	if err != nil {
		Warn.Printf("Unable to Env file: %v", err)
	}
	protocol := os.Getenv("SERVER_PROTOCOL")
	if protocol == "" {
		protocol = "tcp"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8900"
	}
	return ServerConfig{
		Port:     port,
		Protocol: protocol,
	}
}

func main() {
	// Loading environmental variables
	cfg := loadenvvariable()

	listener, err := net.Listen(cfg.Protocol, fmt.Sprintf(":%s", cfg.Port))
	Info.Printf("Snap Basket Protocol: %s, Port: %s", cfg.Protocol, cfg.Port)
	if err != nil {
		panic(err)
	}

	// Database Connection Establishment
	repo.InitMongo()

	// Register gRPC Model Servers
	user_server := &server.UserServer{}
	s := grpc.NewServer()

	// Register Servers explicitly
	pb.RegisterCreateUserModelServiceServer(s, user_server)
	pb.RegisterLoginUserServiceServer(s, user_server)

	Info.Printf("Snap Basket Application Is Live !")
	server_err := s.Serve(listener)
	if server_err != nil {
		Error.Fatalf("Failed To Serve Snap Basket Application: %v", err)
	}
}
