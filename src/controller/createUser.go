package controller

import (
	"fmt"

	"github.com/derafael04/meu-primeiro-projeto-go/src/configuration/validation"
	"github.com/derafael04/meu-primeiro-projeto-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.Context) {

	var userRequest request.UserRequest

	if err := r.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		r.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
