package service

import (
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Deleting user", zap.String("journey", "deleteUser"))
	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "deleteUser"))
		return err
	}
	return nil
}
