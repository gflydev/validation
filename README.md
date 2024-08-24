# Validation

By default, data checking is supported for structs by `ValidateStruct(structData interface{}, msgForTag MsgForTagFunc) (map[string][]string, error)`

### Usage

Install
```bash
go get -u github.com/gflydev/validation@v1.0.0
```

Quick usage
```go
import "github.com/gflydev/validation"

if errorData, err := validation.Check(loginDto, validation.MsgForTag); err != nil {
    return c.BadRequest(errorData)
}
```

Customize error message by yourself
```go
if errorData, err := validation.CheckData(loginDto, func(fe validator.FieldError) string {
    switch fe.Tag() {
    case "required":
        return "This field is required"
    case "gte":
        return fmt.Sprintf("This field is gte %s characters", fe.Param())
    case "email":
        return "Invalid email"
    }

    return fe.Error()
}); err != nil {
    return c.BadRequest(errorData)
}
```

### Message for Tag

Review and add more code in file `message_for_tag.go`.

### Add custom Validator `uuid_rule.go`

```go
package validation

import (
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
)

// UuidRule Custom validation for uuid fields. Use `validate:"uuid"`
type UuidRule string

func (v UuidRule) GetTag() string {
    return string(v)
}

func (v UuidRule) Handler() validator.Func {
    return func(fl validator.FieldLevel) bool {
        field := fl.Field().String()

        if _, err := uuid.Parse(field); err != nil {
            return false
        }
        return true
    }
}
```
Add rule to validator
```go
import "github.com/vinhio/gfly-modules/validation"

// Add rule before validation.Validate()
validation.AddRule(uuidValidator("uuid"))
```
