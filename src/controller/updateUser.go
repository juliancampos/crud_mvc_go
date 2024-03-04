package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Update User - NOT IMPLEMENTED", zap.String("journey", "updateUser"))
}
