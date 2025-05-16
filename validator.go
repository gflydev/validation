package validation

import (
	goerrors "errors"
	"github.com/gflydev/core"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// ===========================================================================================================
//                                       Validator
// ===========================================================================================================

var (
	customValidators []ICustomValidator
)

// ICustomValidator Custom Validator
type ICustomValidator interface {
	GetTag() string
	Handler() validator.Func
}

// AddRule append a custom validator
func AddRule(validatorObj ICustomValidator) {
	customValidators = append(customValidators, validatorObj)
}

// instance A singleton Validator instance.
var instance *validator.Validate

// ValidatorInstance func for create a new validator for model fields.
func ValidatorInstance() *validator.Validate {
	if instance != nil {
		return instance
	}

	// Create a new validator for a Book model.
	instance = validator.New()

	// Custom validation for myType fields. Use `validate:"myType"`
	for _, validatorObj := range customValidators {
		_ = instance.RegisterValidation(validatorObj.GetTag(), validatorObj.Handler())
	}

	// Get json tag value
	instance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return instance
}

// ===========================================================================================================
// 											Validation functions
// ===========================================================================================================

// CheckData verify a data struct type.
func checkData(structData any, msgForTag MsgForTagFunc) (core.Data, error) {
	validatorInstance := ValidatorInstance()
	var out core.Data

	// Validate data
	err := validatorInstance.Struct(structData)
	if err != nil {
		// Determine error type ValidationErrors.
		var ve validator.ValidationErrors
		if goerrors.As(err, &ve) {
			out = make(core.Data, len(ve))
			// Parse error to build a custom message
			for _, fe := range ve {
				// Get list message of specific field
				messages, _ := out[fe.Field()].([]string)
				// Get new message related field
				message := msgForTag(fe)

				// Append a new message and assign to pool
				out[fe.Field()] = append(messages, message)
			}
		}
	}

	return out, err
}

// Check Validate data struct type.
func Check(structData any, msgForTagFunc ...MsgForTagFunc) (core.Data, error) {
	// Default message tag function.
	fn := MsgForTag

	if len(msgForTagFunc) > 0 {
		fn = msgForTagFunc[0]
	}

	return checkData(structData, fn)
}
