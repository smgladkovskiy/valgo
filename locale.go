package valgo

import (
	"sync"
)

type LocaleInterface interface {
	GetCode() LocaleCode
	GetMessageByKey(key string) (string, bool)
	SetMessage(key, message string)
}

type LocaleCode string

type locale struct {
	code     LocaleCode
	messages map[string]string
	mu       sync.RWMutex
}

func (l *locale) SetMessage(key, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.messages[key] = message
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func NewLocale(code LocaleCode, messages map[string]string) LocaleInterface {
	return &locale{code: code, messages: messages}
}

func (l *locale) GetMessageByKey(key string) (string, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	m, exists := l.messages[key]

	return m, exists
}

func (l *locale) GetCode() LocaleCode {
	return l.code
}
