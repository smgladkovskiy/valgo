package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLocalization(t *testing.T) {
	t.Parallel()

	v, err := valgo.New(valgo.WithDefaultLocale(valgo.LocaleCodeEs))
	require.NoError(t, err)

	result := v.Validate().Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, result.ErrorByKey("value_0").Messages(), "Value 0 no puede estar en blanco")

	// Default localization must be persistent
	result = v.Validate().Is(valgo.String(" ").Empty())
	assert.Contains(t, result.ErrorByKey("value_0").Messages(), "Value 0 debe estar vacío")
}

func TestSeparatedLocalization(t *testing.T) {
	t.Parallel()

	v, err := valgo.New(valgo.WithDefaultLocale(valgo.LocaleCodeEn)) // Default localization is English
	require.NoError(t, err)

	result := v.ValidateForLocale(valgo.LocaleCodeEs).Check(valgo.String(" ", "my_value").Not().Blank().Empty())
	assert.Contains(t, result.ErrorByKey("my_value").Messages(), "My value no puede estar en blanco")
	assert.Contains(t, result.ErrorByKey("my_value").Messages(), "My value debe estar vacío")

	// Default localization must not be changed
	v = v.Validate().Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, v.ErrorByKey("value_0").Messages(), "Value 0 can't be blank")
}

func TestAddLocalization(t *testing.T) {
	t.Parallel()

	l := valgo.NewLocale("ee", map[string]string{valgo.ErrorKeyNotBlank: "{{title}} ei tohi olla tühi"})

	v, err := valgo.New(valgo.WithLocale(l, false))
	require.NoError(t, err)

	result := v.ValidateForLocale("ee").Check(valgo.String(" ", "my_value").Not().Blank())
	assert.Contains(t, result.ErrorByKey("my_value").Messages(), "My value ei tohi olla tühi")

	// Default localization must not be changed
	result = v.Validate().Is(valgo.String(" ").Not().Blank())
	assert.Contains(t, result.ErrorByKey("value_0").Messages(), "Value 0 can't be blank")
}
