package test

import (
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestAStringValid(t *testing.T) {
	valgo.ResetMessages()

	v := valgo.Is("a").AString()
	assert.True(t, v.Valid())
	assert.Empty(t, v.ErrorItems())
}

func TestAStringInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		true,
		10,
		10.1,
		&[]int{10},
		[]int{10}} {
		v := valgo.Is(value).AString()

		if assert.NotEmpty(t, v.ErrorItems()) {
			assert.Len(t, v.ErrorItems(), 1)
			assert.Contains(t, v.ErrorItems()[0].Messages, "\"value0\" must be a text")
		}
	}
}
