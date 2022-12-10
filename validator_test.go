package valgo_test

import (
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ValidatorSecretWord struct {
	context *valgo.ValidatorContext
}

func (validator *ValidatorSecretWord) Correct(template ...string) *ValidatorSecretWord {
	validator.context.Add(
		func() bool {
			strVal, ok := validator.context.Value().(string)
			if !ok {
				return false
			}

			check := strVal == "cohesive" || strVal == "stack"

			return check
		},
		"not_valid_secret", template...)

	return validator
}

func (validator *ValidatorSecretWord) Context() *valgo.ValidatorContext {
	return validator.context
}

func TestCustomValidator(t *testing.T) {
	t.Parallel()

	v, err := valgo.New(valgo.WithDefaultLocale(valgo.LocaleCodeEn))
	require.NoError(t, err)

	defaultLocale, err := v.GetDefaultLocale()
	assert.NoError(t, err)

	defaultLocale.SetMessage("not_valid_secret", "{{title}} is invalid.")

	v.Is(SecretWord("loose", "secret").Correct())

	assert.False(t, v.Valid())
	assert.Equal(t, v.ErrorsCount(), 1)
	assert.Len(t, v.ErrorByKey("secret").Messages(), 1)
	assert.Contains(t, v.ErrorByKey("secret").Messages(), "Secret is invalid.")

	v2 := v.ValidateForLocale(valgo.LocaleCodeEn).Is(SecretWord("cohesive", "secret").Correct())
	assert.True(t, v2.Valid())
	assert.Equal(t, v2.ErrorsCount(), 0)
}

func SecretWord(value string, nameAndTitle ...string) *ValidatorSecretWord {
	return &ValidatorSecretWord{context: valgo.NewContext(value, nameAndTitle...)}
}
