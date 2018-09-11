package errorbag

import (
	"errors"
	"strings"
	"sync"
)

type ErrorBag struct {
	sync.Mutex
	errors []error
}

func (e *ErrorBag) Add(err error) {
	if err == nil {
		return
	}

	e.Lock()
	e.errors = append(e.errors, err)
	e.Unlock()
}

func (e *ErrorBag) Error() error {
	e.Lock()
	defer e.Unlock()

	if len(e.errors) == 0 {
		return nil
	}

	var errs []string
	for _, err := range e.errors {
		errs = append(errs, err.Error())
	}

	return errors.New(strings.Join(errs, "\n"))
}
