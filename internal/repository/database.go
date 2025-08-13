package repository

import (
	"context"
	"log"
	"time"

	loaders "github.com/karthikbalasubramani/snap-basket/internal/config"
	logx "github.com/karthikbalasubramani/snap-basket/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

// Global Declarations for logger
var Info = logx.CustomLogger.Info
var Error = logx.CustomLogger.Error
var Warn = logx.CustomLogger.Warn

// Initialize Mongo Database
func InitMongo() {
	cfgDatbaseConfig := loaders.LoadDatabseConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfgDatbaseConfig.Uri))
	if err != nil {
		log.Fatal(err)
	}
	Info.Println("Database Connection Established")
	UserCollection = client.Database(cfgDatbaseConfig.DatabaseName).Collection(cfgDatbaseConfig.UserCollection)
}
