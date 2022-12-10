package valgo

import (
	"errors"
	"fmt"
)

var (
	ErrLocaleDoesNotExist     = errors.New("doesn't exist a registered locale with code")
	ErrLocalesAreNotInitiated = errors.New("locales are not initiated")
)

func localeDoesNotExist(code string) error {
	return fmt.Errorf("%w '%s'", ErrLocaleDoesNotExist, code)
}
