package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorNumberPNot(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	number1 := 2

	v := valgo.Is(valgo.NumberP(&number1).Not().EqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).EqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).EqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 2
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).EqualTo(3))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.NumberP(&myNumber1).EqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Equal(t,
		"Value 0 must be equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPGreaterThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 3

	v = valgo.Is(valgo.NumberP(&number1).GreaterThan(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).GreaterThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPGreaterThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 2
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).GreaterThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	*number1 = 2

	v = valgo.Is(valgo.NumberP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).GreaterThan(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).GreaterThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPGreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 3

	v = valgo.Is(valgo.NumberP(&number1).GreaterOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPGreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 2
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).GreaterOrEqualTo(3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.NumberP(&myNumber1).GreaterOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"3\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPLessThanValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).LessThan(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3

	v = valgo.Is(valgo.NumberP(&myNumber1).LessThan(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPLessThanInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 2
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	*number1 = 3

	v = valgo.Is(valgo.NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).LessThan(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).LessThan(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPLessOrEqualToValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).LessOrEqualTo(2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = valgo.Is(valgo.NumberP(&number1).LessOrEqualTo(3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPLessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 3
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil
	v = valgo.Is(valgo.NumberP(number1).LessOrEqualTo(2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 3
	var myNumber2 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).LessOrEqualTo(myNumber2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPBetweenValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 4

	v = valgo.Is(valgo.NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = 6

	v = valgo.Is(valgo.NumberP(&number1).Between(2, 6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 2
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.NumberP(&myNumber1).Between(myNumber2, myNumber3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPBetweenInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 2
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	*number1 = 7

	v = valgo.Is(valgo.NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).Between(3, 6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2
	var myNumber2 MyNumber = 3
	var myNumber3 MyNumber = 6

	v = valgo.Is(valgo.NumberP(&myNumber1).Between(myNumber2, myNumber3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"3\" and \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPZeroValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 0
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	_number2 := 0.0
	number2 := &_number2

	v = valgo.Is(valgo.NumberP(number2).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber

	v = valgo.Is(valgo.NumberP(&myNumber1).Zero())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPZeroInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 1
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.NumberP(&myNumber1).Zero())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPZeroOrNilValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	_number1 := 0
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber

	v = valgo.Is(valgo.NumberP(&myNumber1).ZeroOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPZeroOrNilInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))
	var v *valgo.Validation

	number1 := 1

	v = valgo.Is(valgo.NumberP(&number1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.NumberP(&myNumber1).ZeroOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be zero",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPPassingValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPPassingInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := 1

	v = valgo.Is(valgo.NumberP(&number1).Passing(func(val *int) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.NumberP(&myNumber1).Passing(func(val *MyNumber) bool {
		return *val == 2
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberPInSliceValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	number1 := 2

	v = valgo.Is(valgo.NumberP(&number1).InSlice([]int{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 2

	v = valgo.Is(valgo.NumberP(&myNumber1).InSlice([]MyNumber{1, 2, 3}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberPInSliceInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	_number1 := 1
	number1 := &_number1

	v = valgo.Is(valgo.NumberP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	number1 = nil

	v = valgo.Is(valgo.NumberP(number1).InSlice([]int{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.NumberP(&myNumber1).InSlice([]MyNumber{2, 3, 4}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorNumberNilIsValid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	var valNumber *int

	v = valgo.Is(valgo.NumberP(valNumber).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyNumber int
	var myNumber1 *MyNumber

	v = valgo.Is(valgo.NumberP(myNumber1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorNumberNilIsInvalid(t *testing.T) {
	t.Parallel()
	require.NoError(t, TearUpTest(t))

	var v *valgo.Validation

	valNumber := 1

	v = valgo.Is(valgo.NumberP(&valNumber).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyNumber int
	var myNumber1 MyNumber = 1

	v = valgo.Is(valgo.NumberP(&myNumber1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])
}
