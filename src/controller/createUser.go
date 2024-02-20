package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/configuration/rest_err"
	"github.com/juliancampos/crud_mvc_go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Some incorret fields, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
