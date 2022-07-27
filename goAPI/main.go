package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"goAPI/configs"
	"goAPI/controllers"
	"goAPI/routes"
	"os"
)

func init() {
	err := os.Mkdir(configs.UsersData, 0777)
	if err != nil {
		return
	}

}

func main() {
	mongoClient := configs.ConnectDB()
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, context.TODO())

	controllers.InitCollection(mongoClient, configs.DbName, configs.ColName)
	router := gin.Default()
	routes.Routes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
