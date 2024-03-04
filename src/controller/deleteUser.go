package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Delete User - NOT IMPLEMENTED", zap.String("journey", "deleteUser"))
}
