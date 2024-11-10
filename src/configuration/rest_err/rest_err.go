package RestErr

import "net/http"

type RestErr struct {
	Mensage string   `json:"mensage"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Mensage string `json:"mensage"`
}

func (e *RestErr) Error() string {
	return e.Mensage
}

func NewError(mensage string, err string, code int, cause []Causes) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     err,
		Code:    code,
		Causes:  cause,
	}
}

func NewBadRequestError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidatorError(mensage string, cause []Causes) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  cause,
	}
}

func NewInternalServerError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewAuthenticationError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "authentication_error",
		Code:    http.StatusUnauthorized,
	}
}

func NewFobbidenError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "fobbiden",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(mensage string) *RestErr {
	return &RestErr{
		Mensage: mensage,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
