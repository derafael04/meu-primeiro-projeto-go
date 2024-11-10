package controller

import (
	"fmt"

	RestErr "github.com/derafael04/meu-primeiro-projeto-go/src/configuration/rest_err"
	"github.com/derafael04/meu-primeiro-projeto-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.Context) {

	var userRequest request.UserRequest

	if err := r.ShouldBindJSON(&userRequest); err != nil {
		restErr := RestErr.NewBadRequestError(fmt.Sprintf("invalid json body: %s", err.Error()))

		r.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
