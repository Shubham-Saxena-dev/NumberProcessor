package customerrors

import (
	"fmt"
	"sync"
)

type ErrorCollector struct {
	Errors map[string][]error
	mutex  sync.Mutex
}

func NewErrorCollector() *ErrorCollector {
	return &ErrorCollector{
		Errors: make(map[string][]error),
	}
}

func (c *ErrorCollector) AddErr(numberType string, err error) {
	if _, ok := c.Errors[numberType]; !ok {
		c.Errors[numberType] = []error{err}
	} else {
		c.Errors[numberType] = append(c.Errors[numberType], err)
	}
}

func (c *ErrorCollector) Error() string {
	msg := "Errors:\n"
	for nType, errors := range c.Errors {
		msg += fmt.Sprintf("Number Type: %s\n", nType)
		for _, e := range errors {
			msg += fmt.Sprintf("\tError: %s\n", e.Error())
		}
	}
	return msg
}
