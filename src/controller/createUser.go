package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"github.com/juliancampos/crud_mvc_go/src/configuration/validation"
	"github.com/juliancampos/crud_mvc_go/src/controller/model/request"
	"github.com/juliancampos/crud_mvc_go/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Create User", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journy", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	print(domainResult)
	c.JSON(http.StatusOK, "User created successfully")
}
