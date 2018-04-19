package valgo

import (
	"fmt"
	"strings"

	"github.com/valyala/fasttemplate"
)

type Validator struct {
	currentValue *Value
	currentTitle string
	currentName  string
	currentValid bool
	currentIndex int
	currentError *Error

	_currentValueAsString *string

	_locale locale
	valid   bool
	errors  []*Error
}

func (validator *Validator) Is(value interface{}) *Validator {

	validator.currentIndex += 1
	validator.currentValue = NewValue(value)
	switch value.(type) {
	case string:
		valueAsString := value.(string)
		validator._currentValueAsString = &valueAsString
	}
	validator.currentValid = true
	validator.currentName = fmt.Sprintf("value%v", validator.currentIndex)
	validator.currentTitle = validator.currentName

	return validator
}

func (validator *Validator) Named(name string) *Validator {
	validator.currentName = name

	return validator
}

func (validator *Validator) Titled(title string) *Validator {
	validator.currentTitle = title

	return validator
}

func (validator *Validator) Valid() bool {
	return validator.valid
}

func (validator *Validator) Passing(
	function func(cv *CustomValidator, t ...string), template ...string) *Validator {

	customValidator := CustomValidator{
		validator: validator,
	}

	if len(template) > 0 {
		function(&customValidator, template[0])
	} else {
		function(&customValidator)
	}
	return validator
}

func (validator *Validator) Errors() []*Error {
	return validator.errors
}

func (validator *Validator) invalidate(
	key string, values map[string]interface{}, templateString []string) {

	validator.valid = false
	var _templateString string
	if len(templateString) > 0 {
		_templateString = templateString[0]
	} else if ts, ok := validator._locale.Messages[key]; ok {
		_templateString = ts
	} else if len(strings.TrimSpace(key)) == 0 {
		_templateString = "ERROR: MISSING MESSAGE KEY OR TEMPLATE STRING!"
	} else {
		_templateString = fmt.Sprintf(
			"ERROR: THERE IS NOT A MESSAGE WITH THE KEY \"%s\"!", key)
	}

	template := fasttemplate.New(_templateString, "{{", "}}")
	message := template.ExecuteString(values)

	if validator.currentError == nil {
		validator.currentError = &Error{
			Name:  validator.currentName,
			Title: validator.currentTitle,
			Value: validator.currentValue,
		}

		validator.currentError.Messages = []string{message}
		validator.currentValid = false
		validator.valid = false

		if validator.errors == nil {
			validator.errors = []*Error{validator.currentError}
		} else {
			validator.errors = append(validator.errors, validator.currentError)
		}
	} else {
		validator.currentError.Messages = append(
			validator.currentError.Messages, message)
	}
}
