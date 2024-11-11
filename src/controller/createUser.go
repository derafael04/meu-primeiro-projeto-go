package controller

import (
	"net/http"

	"github.com/derafael04/meu-primeiro-projeto-go/src/configuration/logger"
	"github.com/derafael04/meu-primeiro-projeto-go/src/configuration/validation"
	"github.com/derafael04/meu-primeiro-projeto-go/src/controller/model/request"
	"github.com/derafael04/meu-primeiro-projeto-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(r *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journey", "create_user"),
	)

	var userRequest request.UserRequest

	if err := r.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "create_user"),
		)
		restErr := validation.ValidateUserError(err)

		r.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		r.JSON(err.Code, err)
		return
	}

	logger.Info("Created user successfully",
		zap.String("journey", "create_user"),
	)
	r.JSON(http.StatusOK, "")

}
