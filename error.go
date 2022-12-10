package valgo

import (
	"encoding/json"
	"fmt"
)

// Error Implementation of the Go error interface in Valgo. The [Validation.Error()]
// method returns a value of this type.
//
// There is a function in this type, [Errors()], that returns a list of errors
// in a [Validation] session.
type Error struct {
	customMarshalJSONFunc func(e *Error) ([]byte, error)
	errors                map[string]ValueErrorInterface
}

// Error Return the error message associated with a Valgo error.
func (e *Error) Error() string {
	count := len(e.errors)
	if count == 1 {
		return "There is 1 error"
	}

	return fmt.Sprintf("There are %v errors", count)
}

// Errors Return a map with each Invalid value error.
func (e *Error) Errors() map[string]ValueErrorInterface {
	return e.errors
}

// MarshalJSON Return the JSON encoding of the validation error messages.
//
// A custom function can be set with [SetMarshalJson()].
func (e *Error) MarshalJSON() ([]byte, error) {
	if e.customMarshalJSONFunc != nil {
		return e.customMarshalJSONFunc(e)
	}

	errors := make(map[string]any)

	for k, v := range e.errors {
		errors[k] = v.Messages()
	}

	return json.Marshal(errors)
}
