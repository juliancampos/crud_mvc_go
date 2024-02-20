package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliancampos/crud_mvc_go/src/controller"
)

func InitiRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
