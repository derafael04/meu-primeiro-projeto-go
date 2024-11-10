package routes

import (
	"github.com/derafael04/meu-primeiro-projeto-go/src/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/upfateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteuser/:userId", controller.DeleteUser)
}
