package test

import (
	"fmt"
	"testing"

	"git.cohesivestack.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
)

func TestAnIntegerValid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		"1",
		"10",
		"-10",
		10,
		-10,
		uint(10),
		uint8(10),
		uint16(10),
		uint32(10),
		uint64(10),
		int(10),
		int8(10),
		int16(10),
		int32(10),
		int64(10)} {
		v := valgo.Is(value).AnInteger()

		assert.True(t, v.Valid())
		assert.Empty(t, v.ErrorItems(), fmt.Sprintf("not assert using %s", value))
	}
}

func TestAnIntegerInvalid(t *testing.T) {
	valgo.ResetMessages()

	for _, value := range []interface{}{
		"10.1",
		"-10.1",
		"",
		" ",
		"a",
		"a10",
		".10",
		"@1",
		[]int{10},
		10.1,
		float32(10.1),
		float64(10.1)} {
		v := valgo.Is(value).AnInteger()

		if assert.NotEmpty(t, v.ErrorItems()) {
			assert.Len(t, v.ErrorItems(), 1)
			assert.Contains(t, v.ErrorItems()[0].Messages, "\"value0\" must be an integer number")
		}
	}
}
