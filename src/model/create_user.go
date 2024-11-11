package model

import (
	"fmt"

	"github.com/derafael04/meu-primeiro-projeto-go/src/configuration/logger"
	RestErr "github.com/derafael04/meu-primeiro-projeto-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *RestErr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "create_user"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
