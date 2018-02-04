package valgo

import (
	"reflect"
)

func equalTo(valueA interface{}, valueB interface{}) bool {
	if isComparableType(valueA) && isComparableType(valueB) && valueA == valueB {
		return true
	}

	// if pass test was not true and one value is nil then just return false
	if valueA == nil || valueB == nil {
		return false
	}

	rvA := reflect.ValueOf(valueA)
	rvB := reflect.ValueOf(valueB)

	if rvA.Kind() == reflect.Ptr {
		valueA = reflect.Indirect(rvA).Interface()
	}

	if rvB.Kind() == reflect.Ptr {
		valueB = reflect.Indirect(rvB).Interface()
	}

	if aNumberType(valueA) && aNumberType(valueB) {
		_valueA, err := getNumberAsFloat64(valueA)
		if err != nil {
			return false
		}
		_valueB, err := getNumberAsFloat64(valueB)
		if err != nil {
			return false
		}
		return float64(_valueA) == float64(_valueB)
	}

	return reflect.DeepEqual(valueA, valueB)
}

func (validator *Validator) EqualTo(value interface{}, template ...string) *Validator {

	if !equalTo(validator.currentValue, value) {
		validator.invalidate("equal_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotEqualTo(value interface{}, template ...string) *Validator {

	if equalTo(validator.currentValue, value) {
		validator.invalidate("not_equal_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
