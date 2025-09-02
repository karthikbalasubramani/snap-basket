package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/karthikbalasubramani/snap-basket/api/pb"
	logx "github.com/karthikbalasubramani/snap-basket/internal/logger"
	models "github.com/karthikbalasubramani/snap-basket/internal/models"
	repo "github.com/karthikbalasubramani/snap-basket/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Loggers
var Info = logx.CustomLogger.Info
var Error = logx.CustomLogger.Error
var Warn = logx.CustomLogger.Warn

// Hashing Password
func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Verifying the hashed password
func VerifyHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create User handler function which updates Database
func HandlerCreateUser(req *pb.CreateUserRequest, db repo.InsertOneInterface) (*pb.CreateUserResponse, error) {
	// Create a context with timeout for DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// fmt.Printf("Incoming Create User Handler: %v", req)
	// Validate input
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name field cannot be empty")
	}
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "Email field cannot be empty")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Password field cannot be empty")
	}
	// Hashing Password using bcrypt
	hashedpassword, err := HashedPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password received: %v", err)
	}

	// Construct model from request
	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedpassword,
		CreatedAt: time.Now(),
	}

	// Insert user into MongoDB collection
	result, err := db.InsertOne(ctx, user)
	if err != nil {
		// Wrap error with gRPC Internal code
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	// Return success response
	return &pb.CreateUserResponse{
		Message: fmt.Sprintf("User created successfully: Name: %s, ID: %s", user.Name, result.InsertedID),
	}, nil
}

// User login handler function includes password verification using bcrypt
func HandlerLoginUser(req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// Create a context with timeout for DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Checl for either name/Email and Password
	if (req.Username == "" || req.Email == "") && req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username/name or email and password are required")
	}

	var user models.User

	// Get User filter with "or" between name and email
	getUserfilter := bson.M{
		"$or": []bson.M{
			{"name": req.Username},
			{"email": req.Email},
		},
	}
	// Get User document and decode it to User model
	err := repo.UserCollection.FindOne(ctx, getUserfilter).Decode(&user)
	if err != nil {
		return &pb.LoginUserResponse{
			Message: fmt.Sprintf("User Not Found: %v", err),
		}, nil
	}

	// Password Verification using bcrypt
	passwordVerification := VerifyHashedPassword(req.Password, user.Password)
	if !passwordVerification {
		return &pb.LoginUserResponse{
			Message: "User Verification Failed, Password Mismatch",
		}, nil
	} else {
		return &pb.LoginUserResponse{
			Message: "Password Verificaion Successful",
		}, nil
	}
}
