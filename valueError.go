package valgo

import (
	"sync"

	"github.com/valyala/fasttemplate"
)

type ValueErrorInterface interface {
	Messages() []string
	IsDirty(bool)
	AddErrorMessage(message string)
	ErrorTemplateByKey(key string) ErrorTemplateInterface
}

// valueError Contains information about each invalid field value returned by the
// [Validation] session.
type valueError struct {
	name           *string
	title          *string
	errorTemplates map[string]ErrorTemplateInterface
	errorMessages  []string
	messages       []string
	dirty          bool
	locale         LocaleInterface
	mu             sync.RWMutex
}

//nolint:ireturn,nolintlint // Interface need to be returned here
func (ve *valueError) ErrorTemplateByKey(key string) ErrorTemplateInterface {
	ve.mu.Lock()

	et, ok := ve.errorTemplates[key]
	if !ok {
		et = newErrorTemplate(key)
		ve.errorTemplates[key] = et
	}

	ve.mu.Unlock()

	return et
}

func (ve *valueError) AddErrorMessage(message string) {
	ve.errorMessages = append(ve.errorMessages, message)
}

func (ve *valueError) IsDirty(b bool) {
	ve.dirty = b
}

// Messages Error messages related to an invalid field value.
func (ve *valueError) Messages() []string {
	if ve.dirty {
		ve.messages = make([]string, 0)
		for _, et := range ve.errorTemplates {
			ve.messages = append(ve.messages, ve.buildMessageFromTemplate(et))
		}

		ve.messages = append(ve.messages, ve.errorMessages...)

		ve.dirty = false
	}

	return ve.messages
}

func (ve *valueError) buildMessageFromTemplate(et ErrorTemplateInterface) string {
	var ts string
	if et.GetTemplate() != nil {
		ts = *et.GetTemplate()
	} else if _ts, ok := ve.locale.GetMessageByKey(et.GetKey()); ok {
		ts = _ts
	} else {
		ts = concatString("ERROR: THERE IS NOT A MESSAGE WITH THE KEY: ", et.GetKey())
	}

	var title string
	if ve.title == nil {
		title = humanizeName(*ve.name)
	} else {
		title = *ve.title
	}

	et.SetParam("name", *ve.name)
	et.SetParam("title", title)

	t := fasttemplate.New(ts, "{{", "}}")

	return t.ExecuteString(et.GetParams())
}
