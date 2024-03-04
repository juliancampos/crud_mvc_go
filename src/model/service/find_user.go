package service

import (
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"github.com/juliancampos/crud_mvc_go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by email", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}
