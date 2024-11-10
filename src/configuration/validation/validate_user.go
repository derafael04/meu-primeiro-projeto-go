package validation

import (
	"encoding/json"
	"errors"

	RestErr "github.com/derafael04/meu-primeiro-projeto-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *RestErr.RestErr {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonError) {
		return RestErr.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []RestErr.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := RestErr.Causes{
				Mensage: err.Translate(transl),
				Field:   err.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return RestErr.NewBadRequestValidatorError("Some fields are invalid", errorsCauses)
	}

	return RestErr.NewBadRequestError("Error trying to convert fields")
}
