// Code generated by Valgo; DO NOT EDIT.
package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
{{ range . }}
func TestValidator{{ .Name }}Not(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	v := valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).Not().EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}EqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).EqualTo(my{{ .Name }}2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}EqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).EqualTo(2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 1
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).EqualTo(my{{ .Name }}2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}GreaterThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(3)).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 3
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).GreaterThan(my{{ .Name }}2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}GreaterThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).GreaterThan(my{{ .Name }}2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}GreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(3)).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).GreaterOrEqualTo(my{{ .Name }}2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}GreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 3

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).GreaterOrEqualTo(my{{ .Name }}2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}LessThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 3

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).LessThan(my{{ .Name }}2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}LessThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(3)).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).LessThan(my{{ .Name }}2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}LessOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).LessOrEqualTo(my{{ .Name }}2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}LessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(3)).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 3
	var my{{ .Name }}2 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).LessOrEqualTo(my{{ .Name }}2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}BetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(4)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(6)).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 2
	var my{{ .Name }}3 My{{ .Name }} = 6

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).Between(my{{ .Name }}2, my{{ .Name }}3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}BetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(7)).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2
	var my{{ .Name }}2 My{{ .Name }} = 3
	var my{{ .Name }}3 My{{ .Name }} = 6

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).Between(my{{ .Name }}2, my{{ .Name }}3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}ZeroValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(0)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 0

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(my{{ .Name }}1)).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}
func TestValidator{{ .Name }}ZeroInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 1

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(my{{ .Name }}1)).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}PassingValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).Passing(func(val {{ .Type }}) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 1

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).Passing(func(val My{{ .Name }}) bool {
		return val == 1
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}PassingInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(1)).Passing(func(val {{ .Type }}) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 1

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).Passing(func(val My{{ .Name }}) bool {
		return val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidator{{ .Name }}InSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(2)).InSlice([]{{ .Type }}{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 2

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).InSlice([]My{{ .Name }}{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidator{{ .Name }}InSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	v = valgo.Is(valgo.{{ .Name }}({{ .Type }}(4)).InSlice([]{{ .Type }}{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type My{{ .Name }} {{ .Type }}
	var my{{ .Name }}1 My{{ .Name }} = 4

	v = valgo.Is(valgo.{{ .Name }}(My{{ .Name }}(my{{ .Name }}1)).InSlice([]My{{ .Name }}{1, 2, 3}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}

{{ end }}