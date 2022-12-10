package valgo

import (
	"sync"
)

type localization struct {
	defaultCode LocaleCode

	mu      sync.RWMutex
	locales map[LocaleCode]LocaleInterface
}

func basicLocales() *localization {
	locales := localization{locales: make(map[LocaleCode]LocaleInterface)}

	locales.SetLocale(&localeEn, true)
	locales.SetLocale(&localeEs, false)

	return &locales
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func (l *localization) GetLocaleByCode(code LocaleCode) (LocaleInterface, error) {
	if l == nil || len(l.locales) == 0 {
		return nil, ErrLocalesAreNotInitiated
	}

	l.mu.RLock()

	loc, ok := l.locales[code]

	l.mu.RUnlock()

	if !ok {
		return nil, ErrLocaleDoesNotExist
	}

	return loc, nil
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func (l *localization) GetDefaultLocale() (LocaleInterface, error) {
	if l == nil {
		return nil, ErrLocalesAreNotInitiated
	}

	return l.GetLocaleByCode(l.defaultCode)
}

func (l *localization) SetLocale(locale LocaleInterface, isDefault bool) {
	l.mu.Lock()
	l.locales[locale.GetCode()] = locale
	l.mu.Unlock()

	if isDefault {
		l.defaultCode = locale.GetCode()
	}
}

func (l *localization) SetDefaultLocaleCode(code LocaleCode) error {
	if _, err := l.GetLocaleByCode(code); err != nil {
		return err
	}

	l.defaultCode = code

	return nil
}
