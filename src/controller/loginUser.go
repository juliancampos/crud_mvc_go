package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/configuration/logger"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {

	logger.Info("Login User - NOT IMPLEMENTED", zap.String("journey", "loginUser"))
	c.JSON(http.StatusOK, "Login User")
}
