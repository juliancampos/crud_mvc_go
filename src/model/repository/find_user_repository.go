package repository

import (
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"github.com/juliancampos/crud_mvc_go/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Finding user by email", zap.String("journey", "findUserByEmail"))
	collection := ur.databaseConnection.Collection("users")

}
