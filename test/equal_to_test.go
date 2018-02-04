package test

import (
	"fmt"
	"testing"

	"github.com/carlosforero/valgo"
	"github.com/stretchr/testify/assert"
)

func TestEqualToValid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	for description, values := range map[string][]interface{}{
		"integers":                []interface{}{1, 1},
		"strings":                 []interface{}{"a", "a"},
		"float integer":           []interface{}{10.0, 10},
		"pointer-integer integer": []interface{}{&_integer, 10},
		"pointer-float float":     []interface{}{&_float, 10.0},
		"pointer-string string":   []interface{}{&_string, "a"},
		"array":                   []interface{}{[]int{10}, []int{10}},
		"pointer-array":           []interface{}{&[]int{10}, &[]int{10}},
		"pointer-array array":     []interface{}{&[]int{10}, []int{10}},
		"map":                   []interface{}{map[string]int{"a": 10}, map[string]int{"a": 10}},
		"pointer-map":           []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 10}},
		"pointer-map map":       []interface{}{&map[string]int{"a": 10}, map[string]int{"a": 10}},
		"struct":                []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
		"pointer-struct":        []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 10}},
		"pointer-struct struct": []interface{}{&MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).EqualTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestEqualToInvalid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	for description, values := range map[string][]interface{}{
		"integers":                []interface{}{1, 2},
		"strings":                 []interface{}{"ab", "a"},
		"string integer":          []interface{}{"1", 1},
		"string float":            []interface{}{"1.0", 1.0},
		"float integer":           []interface{}{10.0, 10.1},
		"pointer-integer integer": []interface{}{&_integer, 11},
		"pointer-integer string":  []interface{}{&_integer, "10.0"},
		"pointer-float float":     []interface{}{&_float, 10.1},
		"pointer-float integer":   []interface{}{&_float, "10.0"},
		"pointer-string string":   []interface{}{&_string, "ab"},
		"array":                   []interface{}{[]int{10}, []int{11}},
		"pointer-array":           []interface{}{&[]int{10}, &[]int{11}},
		"pointer-array array":     []interface{}{&[]int{10}, []int{11}},
		"map":                   []interface{}{map[string]int{"a": 10}, map[string]int{"a": 11}},
		"pointer-map":           []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 11}},
		"pointer-map map":       []interface{}{&map[string]int{"a": 10}, map[string]int{"a": 11}},
		"struct":                []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
		"pointer-struct":        []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 11}},
		"pointer-struct struct": []interface{}{&MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).EqualTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" must be identical to \"%v\"", valueB), msg)
		}
	}
}

func TestNotEqualToValid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	for description, values := range map[string][]interface{}{
		"integers":                []interface{}{1, 2},
		"strings":                 []interface{}{"ab", "a"},
		"string integer":          []interface{}{"1", 1},
		"string float":            []interface{}{"1.0", 1.0},
		"float integer":           []interface{}{10.0, 10.1},
		"pointer-integer integer": []interface{}{&_integer, 11},
		"pointer-integer string":  []interface{}{&_integer, "10.0"},
		"pointer-float float":     []interface{}{&_float, 10.1},
		"pointer-float integer":   []interface{}{&_float, "10.0"},
		"pointer-string string":   []interface{}{&_string, "ab"},
		"array":                   []interface{}{[]int{10}, []int{11}},
		"pointer-array":           []interface{}{&[]int{10}, &[]int{11}},
		"pointer-array array":     []interface{}{&[]int{10}, []int{11}},
		"map":                   []interface{}{map[string]int{"a": 10}, map[string]int{"a": 11}},
		"pointer-map":           []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 11}},
		"pointer-map map":       []interface{}{&map[string]int{"a": 10}, map[string]int{"a": 11}},
		"struct":                []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
		"pointer-struct":        []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 11}},
		"pointer-struct struct": []interface{}{&MyStruct{FieldInt: 10}, MyStruct{FieldInt: 11}},
	} {
		valueA := values[0]
		valueB := values[1]
		msg := fmt.Sprintf("not assert with %s", description)

		v := valgo.Is(valueA).NotEqualTo(valueB)
		assert.True(t, v.Valid(), msg)
		assert.Empty(t, v.Errors(), msg)
	}
}

func TestNotEqualToInvalid(t *testing.T) {
	valgo.ResetMessages()

	_integer := 10
	_float := 10.0
	_string := "a"
	for description, values := range map[string][]interface{}{
		"integers":                []interface{}{1, 1},
		"strings":                 []interface{}{"a", "a"},
		"float integer":           []interface{}{10.0, 10},
		"pointer-integer integer": []interface{}{&_integer, 10},
		"pointer-float float":     []interface{}{&_float, 10.0},
		"pointer-string string":   []interface{}{&_string, "a"},
		"array":                   []interface{}{[]int{10}, []int{10}},
		"pointer-array":           []interface{}{&[]int{10}, &[]int{10}},
		"pointer-array array":     []interface{}{&[]int{10}, []int{10}},
		"map":                   []interface{}{map[string]int{"a": 10}, map[string]int{"a": 10}},
		"pointer-map":           []interface{}{&map[string]int{"a": 10}, &map[string]int{"a": 10}},
		"pointer-map map":       []interface{}{&map[string]int{"a": 10}, map[string]int{"a": 10}},
		"struct":                []interface{}{MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
		"pointer-struct":        []interface{}{&MyStruct{FieldInt: 10}, &MyStruct{FieldInt: 10}},
		"pointer-struct struct": []interface{}{&MyStruct{FieldInt: 10}, MyStruct{FieldInt: 10}},
	} {
		valueA := values[0]
		valueB := values[1]
		v := valgo.Is(valueA).NotEqualTo(valueB)
		msg := fmt.Sprintf("not assert with %s", description)

		assert.False(t, v.Valid())
		if assert.NotEmpty(t, v.Errors(), msg) {
			assert.Len(t, v.Errors(), 1, msg)
			assert.Contains(t, v.Errors()[0].Messages,
				fmt.Sprintf("\"value0\" can't be identical to \"%v\"", valueB), msg)
		}
	}
}
