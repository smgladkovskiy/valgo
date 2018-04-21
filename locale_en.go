package valgo

func init() {
	setDefaultEnglishMessages()
}

func setDefaultEnglishMessages() {
	getLocales()["en"] = locale{
		Messages: map[string]string{
			"blank":                   "\"{{Title}}\" must be blank",
			"not_blank":               "\"{{Title}}\" can't be blank",
			"empty":                   "\"{{Title}}\" must be empty",
			"not_empty":               "\"{{Title}}\" can't be empty",
			"equivalent_to":           "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_equivalent_to":       "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"equal_to":                "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_equal_to":            "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"greater_than":            "\"{{Title}}\" must be greater than \"{{Value}}\"",
			"not_greater_than":        "\"{{Title}}\" can't be greater than \"{{Value}}\"",
			"greater_or_equal_to":     "\"{{Title}}\" must be greater or equal to \"{{Value}}\"",
			"not_greater_or_equal_to": "\"{{Title}}\" can't be greater or equal to \"{{Value}}\"",
			"less_than":               "\"{{Title}}\" must be less than \"{{Value}}\"",
			"not_less_than":           "\"{{Title}}\" can't be less than \"{{Value}}\"",
			"identical_to":            "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_identical_to":        "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"matching_to":             "\"{{Title}}\" must match to \"{{Value}}\"",
			"not_matching_to":         "\"{{Title}}\" can't match to \"{{Value}}\"",
			"a_number":                "\"{{Title}}\" must be a number",
			"a_number_type":           "\"{{Title}}\" must be a number type",
			"an_integer":              "\"{{Title}}\" must be an integer number",
			"an_integer_type":         "\"{{Title}}\" must be an integer type",
			"a_string":                "\"{{Title}}\" must be a text",
		},
	}
}
