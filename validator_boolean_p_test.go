package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorBoolPNot(t *testing.T) {
	t.Parallel()

	bool1 := true

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&bool1).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPEqualToWhenIsValid(t *testing.T) {
	t.Parallel()

	valTrue := true
	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valTrue).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	valFalse := false
	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valFalse).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPEqualToWhenIsInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	valTrue := true

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valTrue).EqualTo(false))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"false\"",
		v.ErrorByKey("value_0").Messages()[0])

	valFalse := false
	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valFalse).EqualTo(true))
	assert.Equal(t,
		"Value 0 must be equal to \"true\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPTrueWhenIsValid(t *testing.T) {
	t.Parallel()

	valTrue := true

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valTrue).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPTrueWhenIsInvalid(t *testing.T) {
	t.Parallel()

	valFalse := false

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valFalse).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).True())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be true",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPFalseWhenIsValid(t *testing.T) {
	t.Parallel()

	_valFalse := false
	valFalse := &_valFalse

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(valFalse).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = false
	mybool1 := &_mybool1

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPFalseWhenIsInvalid(t *testing.T) {
	t.Parallel()

	_valTrue := true
	valTrue := &_valTrue

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(valTrue).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var _mybool1 MyBool = true
	mybool1 := &_mybool1

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(mybool1).False())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be false",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolNilIsValid(t *testing.T) {
	t.Parallel()

	var valTrue *bool

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(valTrue).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 *MyBool

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(mybool1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolNilIsInvalid(t *testing.T) {
	t.Parallel()

	valTrue := true

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valTrue).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPPassingWhenIsValid(t *testing.T) {
	t.Parallel()

	valTrue := true

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valTrue).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPPassingWhenIsInvalid(t *testing.T) {
	t.Parallel()

	valFalse := false

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&valFalse).Passing(func(val *bool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&mybool1).Passing(func(val *MyBool) bool {
		return *val == true
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorBoolPInSliceValid(t *testing.T) {
	t.Parallel()

	boolValue := false

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&boolValue).InSlice([]bool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = false

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&myBool1).InSlice([]MyBool{true, false, true}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorBoolPInSliceInvalid(t *testing.T) {
	t.Parallel()

	_boolValue := true
	boolValue := &_boolValue

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	boolValue = nil

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(boolValue).InSlice([]bool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyBool bool
	var myBool1 MyBool = true

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.BoolP(&myBool1).InSlice([]MyBool{false, false, false}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}
