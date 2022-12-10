package valgo

func WithDefaultLocale(code LocaleCode) func(*Validation) error {
	return func(v *Validation) error {
		return v.localization.SetDefaultLocaleCode(code)
	}
}

func WithLocale(l LocaleInterface, isDefault bool) func(*Validation) error {
	return func(v *Validation) error {
		v.localization.SetLocale(l, isDefault)

		return nil
	}
}

func WithCustomMarshalJSONFunc(customFunc func(e *Error) ([]byte, error)) func(*Validation) error {
	return func(v *Validation) error {
		v.customMarshalJSONFunc = customFunc

		return nil
	}
}
