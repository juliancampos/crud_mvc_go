package repository

import (
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (uc *userRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Deleting user repository - NOT IMPLEMENTED", zap.String("journey", "deleteUser"))
	return nil
}
