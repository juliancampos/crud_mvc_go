package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juliancampos/crud_mvc_go/src/configuration/database/mongodb"
	"github.com/juliancampos/crud_mvc_go/src/controller/routes"
	"golang.org/x/net/context"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitiRoutes(&router.RouterGroup, userController)
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
