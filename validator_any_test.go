package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorAnyNot(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(10).Not().EqualTo(11))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorAnyEqualToValid(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(10).EqualTo(10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	type X struct{ Value int }
	x := X{Value: 10}
	y := X{Value: 10}

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(x).EqualTo(y))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	var a *int
	var b *int

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(a).EqualTo(b))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorAnyEqualToInvalid(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(11).EqualTo(10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	type X struct{ Value int }
	x := X{Value: 10}
	y := X{Value: 11}

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(x).EqualTo(y))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"{11}\"",
		v.ErrorByKey("value_0").Messages()[0])

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(10).EqualTo(nil))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"<nil>\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Both nil but different types
	var a *int
	var b *int64

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(a).EqualTo(b))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be equal to \"<nil>\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorAnyNilValid(t *testing.T) {
	t.Parallel()

	var a *int
	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(a).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	type X struct{}
	var x *X

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(x).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorAnyNilInvalid(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(0).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])

	type X struct{}
	x := X{}

	v, err = valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(&x).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorAnyPassingValid(t *testing.T) {
	t.Parallel()

	valTen := 10

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(valTen).Passing(func(val any) bool {
		return val == 10
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorAnyPassingInvalid(t *testing.T) {
	t.Parallel()

	valTen := 10

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.Any(valTen).Passing(func(val any) bool {
		return val == 9
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}
