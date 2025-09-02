package testing

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/karthikbalasubramani/snap-basket/api/pb"
	"github.com/karthikbalasubramani/snap-basket/internal/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var statusFail = "FAIL"
var statusPass = "PASS"

// Test Create User Success
func TestHandlerCreateUser(t *testing.T) {
	// Test Case Name
	testName := "CreateUserValid"

	mockDB := &mockUserCollection{
		insertOneFn: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil
		},
	}
	req := &pb.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123",
	}

	// Call handler
	resp, err := handlers.HandlerCreateUser(req, mockDB)

	if err != nil {
		RecordTestResult(testName, statusFail, fmt.Sprintf("Unexpected error: %v", err))
	}
	if resp == nil || !strings.Contains(resp.Message, "User created successfully") {
		RecordTestResult(testName, "FAIL", fmt.Sprintf("Unexpected error: %v", err))
	} else {
		RecordTestResult(testName, "PASS", fmt.Sprintf("%v", resp.Message))
	}
}

// Test Create User Missing Field Name
func TestHandlerCreateUser_MissingName(t *testing.T) {
	// Testcase Name
	testCaseName := "CreateUserInvalid_MissingName"
	mockDB := &mockUserCollection{
		insertOneFn: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return nil, nil // Should not be called
		},
	}
	NameInvalidreq := &pb.CreateUserRequest{
		Email:    "Email", // Missing required fields
		Password: "Password",
	}

	resp, err := handlers.HandlerCreateUser(NameInvalidreq, mockDB)
	if resp != nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected nil response, got: %+v", resp))
	}
	if err == nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintln("Expected error for missing fields, got nil"))
	}

	st, ok := status.FromError(err)
	if !ok {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected gRPC status error, got: %v", err))
	}
	if st.Code() == codes.InvalidArgument && st.Message() == "Name field cannot be empty" {
		RecordTestResult(testCaseName, statusPass, fmt.Sprintf("%v", st.Message()))
	} else {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Unexpected error code/message: code=%v message=%q", st.Code(), st.Message()))
	}
}

// Test Create User with misisng field Email
func TestHandlerCreateUser_MissingEmail(t *testing.T) {
	// Testcase Name
	testCaseName := "CreateUserInvalid_MissingEmail"
	mockDB := &mockUserCollection{
		insertOneFn: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return nil, nil // Should not be called
		},
	}
	NameInvalidreq := &pb.CreateUserRequest{
		Name:     "Name", // Missing required fields
		Password: "Password",
	}

	resp, err := handlers.HandlerCreateUser(NameInvalidreq, mockDB)
	if resp != nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected nil response, got: %+v", resp))
	}
	if err == nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintln("Expected error for missing fields, got nil"))
	}

	st, ok := status.FromError(err)
	if !ok {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected gRPC status error, got: %v", err))
	}
	if st.Code() == codes.InvalidArgument && st.Message() == "Email field cannot be empty" {
		RecordTestResult(testCaseName, statusPass, fmt.Sprintf("%v", st.Message()))
	} else {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Unexpected error code/message: code=%v message=%q", st.Code(), st.Message()))
	}
}

func TestHandlerCreateUser_MissingPassword(t *testing.T) {
	// Testcase Name
	testCaseName := "CreateUserInvalid_MissingPassword"
	mockDB := &mockUserCollection{
		insertOneFn: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return nil, nil // Should not be called
		},
	}
	NameInvalidreq := &pb.CreateUserRequest{
		Name:  "Name", // Missing required fields
		Email: "Email",
	}

	resp, err := handlers.HandlerCreateUser(NameInvalidreq, mockDB)
	if resp != nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected nil response, got: %+v", resp))
	}
	if err == nil {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintln("Expected error for missing fields, got nil"))
	}

	st, ok := status.FromError(err)
	if !ok {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Expected gRPC status error, got: %v", err))
	}
	if st.Code() == codes.InvalidArgument && st.Message() == "Password field cannot be empty" {
		RecordTestResult(testCaseName, statusPass, fmt.Sprintf("%v", st.Message()))
	} else {
		RecordTestResult(testCaseName, statusFail, fmt.Sprintf("Unexpected error code/message: code=%v message=%q", st.Code(), st.Message()))
	}
}

func TestPrintTestingStats(t *testing.T) {
	fmt.Println("\n\nTest Results Summary\n----------------------------")
	fmt.Printf("%-40s %-8s %s\n", "Test Case", "Status", "Detail")
	fmt.Printf("%-40s %-8s %s\n", strings.Repeat("-", 39), strings.Repeat("-", 7), strings.Repeat("-", 30))
	for idx, res := range TestResults {
		fmt.Printf("%-40s %-8s %s\n", fmt.Sprintf("%d. %s", idx+1, res.Name), res.Status, res.Detail)
	}
}
