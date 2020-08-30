package repository

import (
	"NewYushinBot/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Setup() error {
	cfg := config.Configs.Mongo
	source := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Auth)

	clientOptions := options.Client().ApplyURI(source)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO()) //nolint
	DB = client.Database(cfg.Name)

	log.Println("Connected to MongoDB!")
	return nil
}
