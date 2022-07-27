package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dbPath    = "mongodb://localhost:27017"
	DbName    = "go_api"
	ColName   = "users"
	UsersData = "usersData"
	//DataToTreat = "dataSET.bin"
)

func ConnectDB() (client *mongo.Client) {
	clientOptions := options.Client().ApplyURI(dbPath)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client

}
