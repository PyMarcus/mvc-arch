package validator

import (
	"encoding/json"
	"errors"

	rest_err "github.com/PyMarcus/mvc-arch/src/configuration/errs"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	translate ut.Translator
	Validate  = validator.New()
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		uni := ut.New(enLocale, enLocale)
		var err error
		translate, _ = uni.GetTranslator("en")
		if err != nil {
			panic("translator initialization failed: " + err.Error())
		}
		if err := en_translations.RegisterDefaultTranslations(val, translate); err != nil {
			panic("failed to register translations: " + err.Error())
		}
	}
}

func ValidateUserError(validation_err error,) *rest_err.Errs {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []*rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(translate),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, &cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
