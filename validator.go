package validation

import (
	goerrors "errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// ===========================================================================================================
// 											Validator
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
func AddRule(validator ICustomValidator) {
	customValidators = append(customValidators, validator)
}

// instance A singleton Validator instance.
var instance *validator.Validate

// ValidatorInstance func for create a new validator for model fields.
func ValidatorInstance() *validator.Validate {
	if instance != nil {
		return instance
	}

	// Create a new validator for a Book model.
	instance := validator.New()

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
func CheckData(structData interface{}, msgForTag MsgForTagFunc) (map[string][]string, error) {
	validatorInstance := ValidatorInstance()
	var out map[string][]string

	// Validate data
	err := validatorInstance.Struct(structData)
	if err != nil {
		// Determine error type ValidationErrors.
		var ve validator.ValidationErrors
		if goerrors.As(err, &ve) {
			out = make(map[string][]string, len(ve))
			// Parse error to build custom message
			for _, fe := range ve {
				messages := out[fe.Field()]
				message := msgForTag(fe)

				out[fe.Field()] = append(messages, message)
			}
		}
	}

	return out, err
}

// Check Validate data struct type.
func Check(structData interface{}, msgForTagFunc ...MsgForTagFunc) (map[string][]string, error) {
	// Default message tag function.
	fn := MsgForTag

	if len(msgForTagFunc) > 0 {
		fn = msgForTagFunc[0]
	}

	return CheckData(structData, fn)
}
