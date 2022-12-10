package valgo

import (
	"fmt"
	"sync"
)

type ErrorTemplateInterface interface {
	GetKey() string
	GetTemplate() *string
	SetTemplate(s string)
	GetParams() map[string]any
	SetParam(key string, param any)
}

type errorTemplate struct {
	key      string
	template *string
	params   map[string]any
	mu       sync.RWMutex
}

func newErrorTemplate(key string) *errorTemplate {
	return &errorTemplate{key: key, params: make(map[string]any)}
}

func (et *errorTemplate) GetParams() map[string]any {
	return et.params
}

func (et *errorTemplate) SetParam(key string, param any) {
	et.mu.Lock()
	et.params[key] = fmt.Sprintf("%v", param)
	et.mu.Unlock()
}

func (et *errorTemplate) GetKey() string {
	return et.key
}

func (et *errorTemplate) GetTemplate() *string {
	return et.template
}

func (et *errorTemplate) SetTemplate(s string) {
	et.template = &s
}
