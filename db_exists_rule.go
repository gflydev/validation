// Package validation provides custom validation rules for database-related checks.
// It includes functionality to validate the existence of values in database tables.
package validation

import (
	"fmt"
	"github.com/gflydev/core/log"
	mb "github.com/gflydev/db"
	"github.com/go-playground/validator/v10"
	qb "github.com/jivegroup/fluentsql"
	"reflect"
	"strings"
)

// TagExists is the default validation tag name for the database existence check.
// It can be used in struct field tags as `validate:"db_exists=table.column"`.
const TagExists ExistsRule = "db_exists"

// The struct to get output of a query
// result represents the output of an EXISTS database query.
// The Exists field indicates whether the queried value was found in the database.
type result struct {
	Exists bool `db:"exists"`
}

// Db is a global database instance obtained from the mb package.
// It provides access to the database connection for executing queries.
var db = mb.Instance()

// Initialize the validator
func init() {
	AddRule(ExistsValidator())
}

// ExistsValidator creates a new DbExistsRule
func ExistsValidator(tag ...string) ExistsRule {
	if len(tag) > 0 {
		return ExistsRule(tag[0])
	}
	return TagExists
}

// ExistsRule is a custom validation rule for checking if a value exists in a database table.
// It supports three validation patterns:
//   - `validate:"db_exists=table.column"` - checks if a single value exists in the table
//   - `validate:"db_exists=all:table.column"` - checks if all array values exist in the table
//   - `validate:"db_exists=one:table.column"` - checks if at least one array value exists in the table
//
// Parameters:
//   - table: the name of the database table to check
//   - column: the name of the column to check against
//
// The rule supports validation of single values and arrays/slices of basic types
// (string, int, uint, float, bool).
type ExistsRule string

func (v ExistsRule) GetTag() string {
	return string(v)
}

func (v ExistsRule) Handler() validator.Func {
	return func(fl validator.FieldLevel) bool {
		// The value of a field and parameter to check
		field := fl.Field()
		// The tag parameter: db:exists=table.column => table.column
		param := fl.Param()

		// Parse the parameter to get table and column
		parts := strings.Split(param, ".")
		if len(parts) != 2 {
			return false
		}

		// Get database table and column name and value to check
		table := parts[0]
		column := parts[1]

		// Handle nil or empty values
		if field.IsZero() {
			return false
		}

		// Get placeholder from default dialect:
		//	- ? Use for MySQL, SQLite
		//  - $ Use for PostgreSQL, SQLite
		placeHolder := qb.DefaultDialect().Placeholder(1)

		// Convert field value to interface{} based on type
		val, ok := convertFieldToInterface(field)
		if !ok {
			return false
		}

		// Single value validation
		out, err := checkValueExists(table, column, placeHolder, val)
		if err != nil {
			log.Tracef("db_exists_rule error: %v", err)

			return false
		}

		return out.Exists
	}
}

// checkValueExists executes a query to check if a value exists in the database
// Returns the query result and any error that occurred
func checkValueExists(table, column, placeHolder string, val interface{}) (result, error) {
	out := result{}
	queryFormat := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = %s)",
		table, column, placeHolder)

	err := db.Raw(queryFormat, val).First(&out)

	return out, err
}

// convertFieldToInterface converts a reflect.Value to an interface{} based on its kind
// Returns the converted value and a boolean indicating success
func convertFieldToInterface(field reflect.Value) (interface{}, bool) {
	switch field.Kind() {
	case reflect.String:
		return field.String(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint(), true
	case reflect.Float32, reflect.Float64:
		return field.Float(), true
	case reflect.Bool:
		return field.Bool(), true
	default:
		return nil, false
	}
}
