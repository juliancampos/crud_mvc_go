package main

import (
	"github.com/juliancampos/crud_mvc_go/src/controller"
	"github.com/juliancampos/crud_mvc_go/src/model/repository"
	"github.com/juliancampos/crud_mvc_go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	rep := repository.NewUserRepository(database)
	service := service.NewUserDomainService(rep)
	return controller.NewUserControllerInterface(service)
}
