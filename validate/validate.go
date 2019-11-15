package validate

import (
	"errors"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var vde *validator.Validate

// VdeInfo .
func VdeInfo(object interface{}) error {

	vde = validator.New()
	err := vde.Struct(object)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				return errors.New(err.StructField() + " is required .")
			case "max":
				return errors.New(err.StructField() + " length exceeds maximum .")
			case "min":
				return errors.New(err.StructField() + " length below minimum .")
			default:
				return errors.New(err.Tag() + "Error message undefined .")
			}
			// fmt.Println(4)
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			// fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println(5)
		}

		// from here you can create your own error messages in whatever language you wish

	}
	return nil
}
