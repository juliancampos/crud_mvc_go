package service

import (
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"github.com/juliancampos/crud_mvc_go/src/model"
	"github.com/juliancampos/crud_mvc_go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
