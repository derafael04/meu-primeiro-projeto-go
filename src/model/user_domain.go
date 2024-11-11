package model

import (
	"crypto/md5"
	"encoding/hex"

	RestErr "github.com/derafael04/meu-primeiro-projeto-go/src/configuration/rest_err"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *RestErr.RestErr
	UpdateUser(string) *RestErr.RestErr
	FindUser(string) (*UserDomain, *RestErr.RestErr)
	DeleteUser(string) *RestErr.RestErr
}

func NewUserDomain(
	email, password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}
