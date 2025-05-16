package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"testing"
)

// MockCustomValidator is a mock implementation of ICustomValidator for testing
type MockCustomValidator struct {
	tag string
}

func (m MockCustomValidator) GetTag() string {
	return m.tag
}

func (m MockCustomValidator) Handler() validator.Func {
	return func(fl validator.FieldLevel) bool {
		// Simple validation that always returns true for testing
		return true
	}
}

// AnotherMockCustomValidator is a second mock implementation of ICustomValidator for testing
type AnotherMockCustomValidator struct {
	tag string
}

func (m AnotherMockCustomValidator) GetTag() string {
	return m.tag
}

func (m AnotherMockCustomValidator) Handler() validator.Func {
	return func(fl validator.FieldLevel) bool {
		// This validator checks if the field value is "valid"
		return fl.Field().String() == "valid"
	}
}

// TestStruct is a test struct with validation tags
type TestStruct struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,gt=0"`
	Optional string `json:"optional"`
}

func TestAddRule(t *testing.T) {
	// Reset customValidators for this test
	customValidators = nil

	// Create mock validators
	mockValidator1 := MockCustomValidator{tag: "mock1"}
	mockValidator2 := AnotherMockCustomValidator{tag: "mock2"}

	// Add the first rule
	AddRule(mockValidator1)

	// Check if the validator was added
	if len(customValidators) != 1 {
		t.Errorf("Expected 1 custom validator, got %d", len(customValidators))
	}

	if customValidators[0].GetTag() != "mock1" {
		t.Errorf("Expected validator tag to be 'mock1', got '%s'", customValidators[0].GetTag())
	}

	// Add the second rule
	AddRule(mockValidator2)

	// Check if both validators are present
	if len(customValidators) != 2 {
		t.Errorf("Expected 2 custom validators, got %d", len(customValidators))
	}

	if customValidators[1].GetTag() != "mock2" {
		t.Errorf("Expected second validator tag to be 'mock2', got '%s'", customValidators[1].GetTag())
	}
}

func TestValidatorInstance(t *testing.T) {
	// Reset the instance for this test
	instance = nil
	customValidators = nil

	// Add custom validators
	mockValidator1 := MockCustomValidator{tag: "mock1"}
	mockValidator2 := AnotherMockCustomValidator{tag: "mock2"}
	AddRule(mockValidator1)
	AddRule(mockValidator2)

	// Get a validator instance
	v := ValidatorInstance()

	// Check if the instance is not nil
	if v == nil {
		t.Error("Expected validator instance to not be nil")
	}

	// Call again to test singleton behavior
	v2 := ValidatorInstance()

	// Check if the same instance is returned
	if v != v2 {
		t.Error("Expected the same validator instance to be returned")
	}

	// Test RegisterTagNameFunc functionality
	type TestJsonTag struct {
		FieldWithTag    string `json:"custom_field_name" validate:"required"`
		FieldWithDash   string `json:"-" validate:"required"`
		FieldWithoutTag string `validate:"required"`
	}

	// Validate a struct with json tags
	err := v.Struct(TestJsonTag{})
	if err == nil {
		t.Error("Expected validation error for empty required fields")
	}

	// Check if validation errors use the json tag names
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		foundCustomField := false
		foundDashField := false

		for _, fe := range validationErrors {
			// Check if the field name matches the json tag
			if fe.Field() == "custom_field_name" {
				foundCustomField = true
			}

			// Check if the field with json:"-" is included
			if fe.Field() == "FieldWithDash" {
				foundDashField = true
			}
		}

		// The field name in the error should match the json tag
		if !foundCustomField {
			t.Errorf("Expected to find field with name 'custom_field_name'")
		}

		// Fields with json:"-" should still be validated
		if !foundDashField {
			t.Errorf("Expected to find field with name 'FieldWithDash'")
		}
	}

	// Test custom validators
	type TestCustomValidation struct {
		Field1 string `validate:"mock1"`
		Field2 string `validate:"mock2"`
	}

	// Test the first validator (always returns true)
	err = v.Struct(TestCustomValidation{Field1: "anything", Field2: "invalid"})
	if err == nil {
		t.Error("Expected validation error for Field2")
	}

	// Test the second validator (checks if value is "valid")
	err = v.Struct(TestCustomValidation{Field1: "anything", Field2: "valid"})
	if err != nil {
		t.Errorf("Expected no validation error when Field2 is 'valid', got %v", err)
	}
}

func TestCheckData(t *testing.T) {
	// Reset the instance for this test
	instance = nil

	tests := []struct {
		name          string
		data          interface{}
		expectedError bool
		expectedKeys  []string
	}{
		{
			name: "Valid data",
			data: TestStruct{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   30,
			},
			expectedError: false,
			expectedKeys:  nil,
		},
		{
			name: "Missing required field",
			data: TestStruct{
				Email: "john@example.com",
				Age:   30,
			},
			expectedError: true,
			expectedKeys:  []string{"name"},
		},
		{
			name: "Invalid email",
			data: TestStruct{
				Name:  "John Doe",
				Email: "invalid-email",
				Age:   30,
			},
			expectedError: true,
			expectedKeys:  []string{"email"},
		},
		{
			name: "Invalid age",
			data: TestStruct{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   0,
			},
			expectedError: true,
			expectedKeys:  []string{"age"},
		},
		{
			name: "Multiple errors",
			data: TestStruct{
				Age: 0,
			},
			expectedError: true,
			expectedKeys:  []string{"name", "email", "age"},
		},
		{
			name:          "Non-struct data",
			data:          "not a struct",
			expectedError: true,
			expectedKeys:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := checkData(tt.data, MsgForTag)

			// Check if error matches expectation
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err != nil)
			}

			// If we expect an error, check if the data contains the expected keys
			if tt.expectedError {
				for _, key := range tt.expectedKeys {
					if _, ok := data[key]; !ok {
						t.Errorf("Expected key '%s' in error data, but it was not found", key)
					}
				}
			}
		})
	}
}

func TestCheck(t *testing.T) {
	// Reset the instance for this test
	instance = nil

	// Test with default MsgForTag function
	t.Run("Default MsgForTag", func(t *testing.T) {
		data := TestStruct{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		}

		result, err := Check(data)

		if err != nil {
			t.Errorf("Expected no error for valid data, got: %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty result for valid data, got: %v", result)
		}
	})

	// Test with custom message function
	t.Run("Custom message function", func(t *testing.T) {
		data := TestStruct{
			// Missing required fields to trigger validation errors
		}

		customMsgFunc := func(fe validator.FieldError) string {
			return "custom error message"
		}

		result, err := Check(data, customMsgFunc)

		if err == nil {
			t.Error("Expected error for invalid data, got nil")
		}

		// Check if all error messages are the custom one
		for field, messages := range result {
			msgArray, ok := messages.([]string)
			if !ok {
				t.Errorf("Expected messages to be []string for field %s", field)
				continue
			}

			for _, msg := range msgArray {
				if msg != "custom error message" {
					t.Errorf("Expected custom error message, got: %s", msg)
				}
			}
		}
	})

	// Test with invalid struct
	t.Run("Invalid struct", func(t *testing.T) {
		// Using a non-struct value should return an error but not panic
		data := "not a struct"

		_, err := Check(data)

		if err == nil {
			t.Error("Expected error for non-struct data, got nil")
		}

		// For non-struct data, we just check that the function doesn't panic
		// and returns an error
	})
}
