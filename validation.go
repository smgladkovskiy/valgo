package valgo

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Validation The [Validation] session in Valgo is the main structure for validating one or
// more values. It is called Validation in code.
//
// A [Validation] session will contain one or more Validators, where each [Validator]
// will have the responsibility to validate a value with one or more rules.
//
// There are multiple functions to create a [Validation] session, depending on the
// requirements:
//
//   - [New]()
//   - [Is](...)
//   - [In](...)
//   - [InRow](...)
//   - [Check](...)
//   - [AddErrorMessage](...)
//
// the function [Is](...) is likely to be the most frequently used function in your
// validations. When [Is](...) is called, the function creates a validation and
// receives a validator at the same time.
type Validation struct {
	localization          *localization
	customMarshalJSONFunc func(e *Error) ([]byte, error)

	valid        bool
	currentIndex int
	errors       map[string]ValueErrorInterface

	mu sync.RWMutex
}

func (v *Validation) Validate() *Validation {
	return v.clear()
}

func (v *Validation) ValidateForLocale(code LocaleCode) *Validation {
	newV := v.clone()
	newV.clear()

	//nolint:errcheck // suppose that it will not be any error here
	_ = newV.localization.SetDefaultLocaleCode(code)

	return newV
}

// Is Add a field validator to a [Validation] session.
func (v *Validation) Is(vr Validator) *Validation {
	return vr.Context().validateIs(v)
}

// Check Add a field validator to a [Validation] session. But unlike [Is()] the
// field validator is not short-circuited.
func (v *Validation) Check(vr Validator) *Validation {
	return vr.Context().validateCheck(v)
}

// Valid A [Validation] session provides this function which returns either true if
// all their validators are valid or false if any one of them is invalid.
//
// In the following example, even though the [Validator] for age is valid, the
// [Validator] for status is invalid, making the entire Validator session
// invalid.
func (v *Validation) Valid() bool {
	return v.valid
}

// In Add a map namespace to a [Validation] session.
func (v *Validation) In(name string, _validation *Validation) *Validation {
	return v.merge(name, _validation)
}

// InRow Add an indexed namespace to a [Validation] session.
func (v *Validation) InRow(name string, index int, _validation *Validation) *Validation {
	return v.merge(fmt.Sprintf("%s[%v]", name, index), _validation)
}

// Merge Using [Merge](...) you can merge two [Validation] sessions. When two
// validations are merged, errors with the same value name will be merged. It is
// useful for reusing validation logic.
//
// The following example merges the [Validation] session returned by the
// validatePreStatus function. Since both [Validation] sessions validate a value
// with the name status, the error returned will return two error messages, and
// without duplicate the Not().Blank() error message rule.
func (v *Validation) Merge(_validation *Validation) *Validation {
	return v.merge("", _validation)
}

//nolint:gocognit // by initial design. should be refactored to be simplified
func (v *Validation) merge(prefix string, _validation *Validation) *Validation {
	var _prefix string
	if len(strings.TrimSpace(prefix)) > 0 {
		_prefix = concatString(prefix, ".")
	}

LOOP1:
	for _field, _err := range _validation.errors {
		for field, err := range v.errors {
			if concatString(_prefix, _field) == field {
			LOOP2:
				for _, _errMsg := range _err.Messages() {
					for _, errMsg := range err.Messages() {
						if _errMsg == errMsg {
							continue LOOP2
						}
					}
					v.AddErrorMessage(concatString(_prefix, _field), _errMsg)
				}

				continue LOOP1
			}
		}

		for _, _errMsg := range _err.Messages() {
			v.AddErrorMessage(concatString(_prefix, _field), _errMsg)
		}
	}

	return v
}

// AddErrorMessage Add an error message to the [Validation] session without executing a field
// validator. By adding this error message, the [Validation] session will be
// marked as invalid.
func (v *Validation) AddErrorMessage(name string, message string) *Validation {
	if v.errors == nil {
		v.errors = make(map[string]ValueErrorInterface)
	}

	v.valid = false

	ev := v.getOrCreateValueError(name)
	ev.AddErrorMessage(message)

	return v
}

func (v *Validation) invalidate(name *string, fragment *validatorFragment) {
	if v.errors == nil {
		v.errors = make(map[string]ValueErrorInterface)
	}

	v.valid = false

	var _name string
	if name == nil {
		_name = concatString("value_", strconv.Itoa(v.currentIndex-1))
	} else {
		_name = *name
	}

	ev := v.getOrCreateValueError(_name)

	errorKey := fragment.errorKey

	if !fragment.boolOperation {
		errorKey = concatString("not_", errorKey)
	}

	et := ev.ErrorTemplateByKey(errorKey)

	if len(fragment.template) > 0 {
		et.SetTemplate(fragment.template[0])
	}

	for k, p := range fragment.templateParams {
		et.SetParam(k, p)
	}
}

// ErrorsCount Return length of the errors map.
func (v *Validation) ErrorsCount() int {
	return len(v.errors)
}

// ErrorByKey Return a map with the information for each invalid field validator in the
// [Validation] session.
//
//nolint:ireturn,nolintlint // Interface need to be returned here
func (v *Validation) ErrorByKey(key string) ValueErrorInterface {
	v.mu.RLock()
	err := v.errors[key]
	v.mu.RUnlock()

	return err
}

// Error Return a map with the information for each invalid field validator in the
// [Validation] session.
func (v *Validation) Error() error {
	if v.valid {
		return nil
	}

	return &Error{customMarshalJSONFunc: v.customMarshalJSONFunc, errors: v.errors}
}

// IsValid Return true if a specific field validator is valid.
func (v *Validation) IsValid(name string) bool {
	if _, isNotValid := v.errors[name]; isNotValid {
		return false
	}

	return true
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func (v *Validation) getOrCreateValueError(name string) ValueErrorInterface {
	v.mu.Lock()
	if _, ok := v.errors[name]; !ok {
		l, _ := v.localization.GetDefaultLocale() //nolint:errcheck // We are pretty sure that it will not be any error here
		v.errors[name] = &valueError{
			name:           &name,
			errorTemplates: map[string]ErrorTemplateInterface{},
			errorMessages:  []string{},
			locale:         l,
		}
	}
	v.mu.Unlock()

	ev := v.errors[name]
	ev.IsDirty(true)

	return ev
}

//nolint:govet // The linter is deliberately ignored with a full understanding of how mutex works in this case.
func (v *Validation) clone() *Validation {
	v.mu.Lock()

	l := *v.localization
	newV := *v
	newV.localization = &l

	v.mu.Unlock()

	newV.mu = sync.RWMutex{}

	return &newV
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func (v *Validation) GetDefaultLocale() (LocaleInterface, error) {
	return v.localization.GetDefaultLocale()
}

func (v *Validation) clear() *Validation {
	v.mu.Lock()

	v.valid = true
	v.currentIndex = 0
	v.errors = nil

	v.mu.Unlock()

	return v
}
