// Code generated by Valgo; DO NOT EDIT.
package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
{{ range . }}
func TestValidator{{ .Name }}PNot(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	number1 := {{ .Type }}(2)

	v := valgo.Is(valgo.{{ .Name }}P(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PGreaterThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(3)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PGreaterThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = 2

	v = valgo.Is(valgo.{{ .Name }}P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PGreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(3)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PGreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PLessThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PLessThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = {{ .Type }}(3)

	v = valgo.Is(valgo.{{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PLessOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.{{ .Name }}P(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PLessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(3)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil
	v = valgo.Is(valgo.{{ .Name }}P(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PBetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(4)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = {{ .Type }}(6)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PBetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(2)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	*number1 = {{ .Type }}(7)

	v = valgo.Is(valgo.{{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PZeroValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(0)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 0

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PZeroInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(1)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PZeroOrNilValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := {{ .Type }}(0)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 0

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PZeroOrNilInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := {{ .Type }}(1)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PPassingValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).Passing(func(val *{{ .Type }}) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PPassingInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := {{ .Type }}(1)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).Passing(func(val *{{ .Type }}) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := {{ .Type }}(2)

	v = valgo.Is(valgo.{{ .Name }}P(&number1).InSlice([]{{ .Type }}{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 2

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}PInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	_number1 := {{ .Type }}(1)
	number1 := &_number1

	v = valgo.Is(valgo.{{ .Name }}P(number1).InSlice([]{{ .Type }}{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.{{ .Name }}P(number1).InSlice([]{{ .Type }}{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PNilIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	var valNumber *{{ .Type }}

	v = valgo.Is(valgo.{{ .Name }}P(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 *MyNumber

	v = valgo.Is(valgo.{{ .Name }}P(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PNilIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valNumber := {{ .Type }}(1)

	v = valgo.Is(valgo.{{ .Name }}P(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyNumber {{ .Type }}
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.{{ .Name }}P(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.Errors()["value_0"].Messages()[0])
}

{{ end }}