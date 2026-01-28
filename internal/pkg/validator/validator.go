package validator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/pkg/errors"
)

type ErrorValidate struct {
	errorValues map[string]string
}

func (er *ErrorValidate) Error() string {
	var writeString strings.Builder

	if er.errorValues != nil {
		return "validation failed"
	}

	for k, v := range er.errorValues {
		fmt.Fprintf(&writeString, "%s: %s \n", k, v)
	}

	return writeString.String()
}

type ValidatePkg interface {
	Validate(data any) error
	ErrorMap(err error, lang string) map[string]string
}

type validatePkg struct {
	validate *validator.Validate
	enTrans  ut.Translator
	idTrans  ut.Translator
}

var (
	once     sync.Once
	instance *validatePkg
)

func NewValidation() ValidatePkg {
	once.Do(func() {
		validate := validator.New(validator.WithRequiredStructEnabled())

		en := en.New()
		id := id.New()
		univ := ut.New(en, en, id)

		enTrans, _ := univ.GetTranslator("en")
		idTrans, _ := univ.GetTranslator("id")

		en_translations.RegisterDefaultTranslations(validate, enTrans)
		id_translations.RegisterDefaultTranslations(validate, idTrans)

		instance = &validatePkg{
			validate: validate,
			enTrans:  enTrans,
			idTrans:  idTrans,
		}

		// Daftarkan custom validation & translation sekali saja
		instance.registerValidation()
		instance.registerTranslation()
	})

	return instance
}

// ErrorMap converts validation error to a map[string]string
func (v *validatePkg) ErrorMap(err error, lang string) map[string]string {
	if err == nil {
		return nil
	}

	if errs, ok := err.(*ErrorValidate); ok {
		return errs.errorValues
	}

	trans := v.enTrans

	switch lang {
	case "id", "ID", "id-ID", "id_ID":
		lang = "id"
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		return validationErrs.Translate(trans)
	}

	return nil
}

/*
register custom validation
*/
func (v *validatePkg) registerValidation() {

}

/*
register custom translate
*/
func (v *validatePkg) registerTranslation() {
	/*
		register your translate "en" of validation here
	*/

	/*
		register your translate "id" of validation here
	*/
}

// validate struct and return error map
func (v *validatePkg) ValidateWithLang(data any, lang string) error {
	trans := v.enTrans

	if lang == "id" {
		trans = v.idTrans
	}
	err := v.validate.Struct(data)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return &ErrorValidate{
				errorValues: errs.Translate(trans),
			}
		}

		return errors.Wrap(err, "validation failed")
	}

	return nil
}

// only validate and return error
func (v *validatePkg) Validate(data any) error {
	return v.ValidateWithLang(data, "en")
}
