// Package valgo Valgo is a type-safe, expressive, and extensible validator library for
// Golang. Valgo is built with generics, so Go 1.18 or higher is required.
//
// Valgo differs from other Golang validation libraries in that the rules are
// written in functions and not in struct tags. This allows greater flexibility
// and freedom when it comes to where and how data is validated.
//
// Additionally, Valgo supports customizing and localizing validation messages.
package valgo

// New This function allows you to create a new Validation session without a
// Validator. This is useful for conditional validation or reusing validation
// logic.
//
// The following example conditionally adds a validator rule for the month_day
// value.
func New(options ...func(*Validation) error) (*Validation, error) {
	v := &Validation{
		valid:        true,
		localization: basicLocales(),
	}

	for _, o := range options {
		if err := o(v); err != nil {
			return nil, err
		}
	}

	return v, nil
}

// Is The [Is](...) function allows you to pass a [Validator] with the value and
// the rules for validating it. At the same time, create a [Validation] session,
// which lets you add more Validators in order to verify more values.
//
// As shown in the following example, we are passing to the function [Is](...)
// the [Validator] for the full_name value. The function returns a [Validation]
// session that allows us to add more Validators to validate more values; in the
// example case the values age and status:.
func Is(v Validator) *Validation {
	nv, _ := New() //nolint:errcheck // suppose that it will not be any error here

	return nv.Is(v)
}

// In The [In](...) function executes one or more validators in a namespace, so the
// value names in the error result are prefixed with this namespace. This is
// useful for validating nested structures.
//
// In the following example we are validating the Person and the nested
// Address structure. We can distinguish the errors of the nested Address
// structure in the error results.
func In(name string, v *Validation) *Validation {
	nv, _ := New() //nolint:errcheck // suppose that it will not be any error here

	return nv.In(name, v)
}

// InRow The [InRow](...) function executes one or more validators in a namespace
// similar to the [In](...) function, but with indexed namespace. So, the value
// names in the error result are prefixed with this indexed namespace. It is
// useful for validating nested lists in structures.
//
// In the following example we validate the Person and the nested list
// Addresses. The error results can distinguish the errors of the nested list
// Addresses.
func InRow(name string, index int, v *Validation) *Validation {
	nv, _ := New() //nolint:errcheck // suppose that it will not be any error here

	return nv.InRow(name, index, v)
}
