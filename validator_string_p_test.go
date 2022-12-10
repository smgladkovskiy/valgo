package valgo_test

import (
	"regexp"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorStringPNot(t *testing.T) {
	t.Parallel()

	text1 := "text1"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).Not().EqualTo("text2"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPEqualToValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "text"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).EqualTo("text"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"
	var myString2 MyString = "text"

	v.Validate().Is(valgo.StringP(&myString1).EqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPEqualToInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "text1"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.ErrorsCount())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).EqualTo("text2"))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.ErrorsCount())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"
	var myString2 MyString = "text2"

	v.Validate().Is(valgo.StringP(&myString1).EqualTo(myString2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.ErrorsCount())
	assert.Equal(t,
		"Value 0 must be equal to \"text2\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPGreaterThanValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "ab"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).GreaterThan("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).GreaterThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPGreaterThanInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "aa"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).GreaterThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = "aa"

	v.Validate().Is(valgo.StringP(text1).GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).GreaterThan("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).GreaterThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPGreaterOrEqualToValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "aa"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "ab"

	v.Validate().Is(valgo.StringP(&text1).GreaterOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).GreaterOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPGreaterOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "aa"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).GreaterOrEqualTo("ab"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v.Validate().Is(valgo.StringP(&myString1).GreaterOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be greater than or equal to \"ab\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPLessThanValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "aa"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).LessThan("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"

	v.Validate().Is(valgo.StringP(&myString1).LessThan(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPLessThanInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "aa"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = "ab"

	v.Validate().Is(valgo.StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).LessThan("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).LessThan(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPLessOrEqualToValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "aa"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).LessOrEqualTo("aa"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	v.Validate().Is(valgo.StringP(&text1).LessOrEqualTo("ab"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).LessOrEqualTo(myString2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPLessOrEqualToInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "ab"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).LessOrEqualTo("aa"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "ab"
	var myString2 MyString = "aa"

	v.Validate().Is(valgo.StringP(&myString1).LessOrEqualTo(myString2))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be less than or equal to \"aa\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPBetweenValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "aa"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "ac"

	v.Validate().Is(valgo.StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "ae"

	v.Validate().Is(valgo.StringP(&text1).Between("aa", "ae"))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "aa"
	var myString3 MyString = "ae"

	v.Validate().Is(valgo.StringP(&myString1).Between(myString2, myString3))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPBetweenInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "aa"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = "af"

	v.Validate().Is(valgo.StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).Between("ab", "ae"))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "aa"
	var myString2 MyString = "ab"
	var myString3 MyString = "ae"

	v.Validate().Is(valgo.StringP(&myString1).Between(myString2, myString3))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be between \"ab\" and \"ae\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPEmptyValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := ""
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString

	v.Validate().Is(valgo.StringP(&myString1).Empty())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPEmptyInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "a"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = " "

	v.Validate().Is(valgo.StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v.Validate().Is(valgo.StringP(&myString1).Empty())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPEmptyOrNilValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := ""
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString

	v.Validate().Is(valgo.StringP(&myString1).EmptyOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPEmptyOrNilInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "a"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = " "

	v.Validate().Is(valgo.StringP(&text1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v.Validate().Is(valgo.StringP(&myString1).EmptyOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be empty",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPBlankValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := ""
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	*text1 = " "

	v.Validate().Is(valgo.StringP(text1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v.Validate().Is(valgo.StringP(&myString1).Blank())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPBlankInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := "a"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v.Validate().Is(valgo.StringP(&myString1).Blank())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPBlankOrNilValid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	_text1 := ""
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	*text1 = " "

	v.Validate().Is(valgo.StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = " "

	v.Validate().Is(valgo.StringP(&myString1).BlankOrNil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPBlankOrNilInvalid(t *testing.T) {
	t.Parallel()
	var v *valgo.Validation

	text1 := "a"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).BlankOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "a"

	v.Validate().Is(valgo.StringP(&myString1).BlankOrNil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be blank",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPPassingValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "text"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).Passing(func(val *string) bool {
		return *val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v.Validate().Is(valgo.StringP(&myString1).Passing(func(val *MyString) bool {
		return *val == "text"
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPPassingInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "text1"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).Passing(func(val *string) bool {
		return *val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text1"

	v.Validate().Is(valgo.StringP(&myString1).Passing(func(val *MyString) bool {
		return *val == "text2"
	}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPInSliceValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "up"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).InSlice([]string{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v.Validate().Is(valgo.StringP(&myString1).InSlice([]MyString{"down", "up", "paused"}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPInSliceInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	_text1 := "up"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).InSlice([]string{"down", "idle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "up"

	v.Validate().Is(valgo.StringP(&myString1).InSlice([]MyString{"down", "indle", "paused"}))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 is not valid",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPMatchingToValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	regex := regexp.MustCompile("pre-.+")

	text1 := "pre-approved"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "pre-approved"

	v.Validate().Is(valgo.StringP(&myString1).MatchingTo(regex))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPMatchingToInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	regex := regexp.MustCompile("pre-.+")

	_text1 := "pre_approved"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "pre_approved"

	v.Validate().Is(valgo.StringP(&myString1).MatchingTo(regex))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must match to \"pre-.+\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPMaxLengthValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "123456"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "12345"

	v.Validate().Is(valgo.StringP(&text1).MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v.Validate().Is(valgo.StringP(&myString1).MaxLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPMaxLengthInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	_text1 := "1234567"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).MaxLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).MaxLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "1234567"

	v.Validate().Is(valgo.StringP(&myString1).MaxLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length longer than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPMinLengthValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "123456"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "1234567"

	v.Validate().Is(valgo.StringP(&text1).MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v.Validate().Is(valgo.StringP(&myString1).MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	myString1 = "1234567"

	v.Validate().Is(valgo.StringP(&myString1).MinLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPMinLengthInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	_text1 := "12345"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v.Validate().Is(valgo.StringP(&myString1).MinLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must not have a length shorter than \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPLengthValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "123456"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).OfLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v.Validate().Is(valgo.StringP(&myString1).OfLength(6))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPLengthInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	_text1 := "12345"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = "1234567"

	v.Validate().Is(valgo.StringP(text1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v.Validate().Is(valgo.StringP(&myString1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])

	myString1 = "1234567"

	v.Validate().Is(valgo.StringP(&myString1).OfLength(6))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length equal to \"6\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringPLengthBetweenValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	text1 := "123456"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&text1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "12345678"

	v.Validate().Is(valgo.StringP(&text1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	text1 = "1234567890"

	v.Validate().Is(valgo.StringP(&text1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 MyString = "123456"

	v.Validate().Is(valgo.StringP(&myString1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	myString1 = "12345678"

	v.Validate().Is(valgo.StringP(&myString1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	myString1 = "1234567890"

	v.Validate().Is(valgo.StringP(&myString1).OfLengthBetween(6, 10))
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringPLengthBetweenInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	_text1 := "12345"
	text1 := &_text1

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(text1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	*text1 = "12345678901"

	v.Validate().Is(valgo.StringP(text1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	text1 = nil

	v.Validate().Is(valgo.StringP(text1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "12345"

	v.Validate().Is(valgo.StringP(&myString1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])

	myString1 = "12345678901"

	v.Validate().Is(valgo.StringP(&myString1).OfLengthBetween(6, 10))
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must have a length between \"6\" and \"10\"",
		v.ErrorByKey("value_0").Messages()[0])
}

func TestValidatorStringNilIsValid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	var valString *string

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(valString).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())

	// Custom Type
	type MyString string
	var myString1 *MyString

	v.Validate().Is(valgo.StringP(myString1).Nil())
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorsCount())
}

func TestValidatorStringNilIsInvalid(t *testing.T) {
	t.Parallel()

	var v *valgo.Validation

	valString := "text"

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.StringP(&valString).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])

	// Custom Type
	type MyString string
	var myString1 MyString = "text"

	v.Validate().Is(valgo.StringP(&myString1).Nil())
	assert.False(t, v.Valid())
	assert.Equal(t,
		"Value 0 must be nil",
		v.ErrorByKey("value_0").Messages()[0])
}
