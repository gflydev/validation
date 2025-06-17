package validation

import (
	"fmt"
	mb "github.com/gflydev/db"
	"github.com/go-playground/validator/v10"
	qb "github.com/jivegroup/fluentsql"
	"reflect"
	"strings"
)

const TagExists ExistsRule = "db_exists"

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

// ExistsRule Custom validation for checking if a value exists in a database table
// Use `validate:"db:exists:table=column"` where:
// - table: the name of the table to check
// - column: the name of the column to check against
type ExistsRule string

func (v ExistsRule) GetTag() string {
	return string(v)
}

func (v ExistsRule) Handler() validator.Func {
	return func(fl validator.FieldLevel) bool {
		db := mb.Instance()

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

		// Convert field value to string based on type
		var value interface{}
		switch field.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = field.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = field.Uint()
		case reflect.Float32, reflect.Float64:
			value = field.Float()
		case reflect.Bool:
			value = field.Bool()
		default:
			return false
		}

		// The struct to get output of query
		type result struct {
			Exists bool `db:"exists"`
		}

		// Get placeholder from default dialect:
		//	- ? Use for MySQL, SQLite
		//  - $ Use for PostgreSQL, SQLite
		placeHolder := qb.DefaultDialect().Placeholder(1)

		out := result{}
		queryFormat := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = %s)",
			table, column, placeHolder)

		if err := db.Raw(queryFormat, value).First(&out); err != nil {
			return false
		}

		return out.Exists
	}
}
