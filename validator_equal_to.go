package valgo

import (
	"reflect"
)

func (valueA *Value) IsEqualTo(value interface{}) bool {
	valueB := NewValue(value)
	if valueA.IsComparableType() && valueB.IsComparableType() && valueA.absolute == valueB.absolute {
		return true
	}

	// if previous test was not true and one value is nil then just return false
	if valueA.absolute == nil || valueB.absolute == nil {
		return false
	}

	if (valueA.IsString() && valueB.IsNumberType()) ||
		(valueB.IsString() && valueA.IsNumberType()) ||
		(valueB.IsNumberType() && valueA.IsNumberType()) {
		return valueA.AsFloat64() == valueB.AsFloat64()
	}

	return reflect.DeepEqual(valueA.absolute, valueB.absolute)
}

func (validator *Validator) EqualTo(value interface{}, template ...string) *Validator {
	if !validator.currentValue.IsEqualTo(value) {
		validator.invalidate("equivalent_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}

func (validator *Validator) NotEqualTo(value interface{}, template ...string) *Validator {
	if validator.currentValue.IsEqualTo(value) {
		validator.invalidate("not_equivalent_to",
			map[string]interface{}{
				"Title": validator.currentTitle,
				"Value": convertToString(value)}, template)
	}
	return validator
}
