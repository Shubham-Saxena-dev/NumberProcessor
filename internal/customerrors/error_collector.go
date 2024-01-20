package customerrors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ErrorHandler struct {
	Errors map[string][]error
	mutex  sync.Mutex
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		Errors: make(map[string][]error),
	}
}

func (c *ErrorHandler) AddErr(numberType string, err error) *ErrorHandler {
	if _, ok := c.Errors[numberType]; !ok {
		c.Errors[numberType] = []error{err}
	} else {
		c.Errors[numberType] = append(c.Errors[numberType], err)
	}
	return c
}

func (c *ErrorHandler) GetErrorsFromNumberType(numberType string) []error {
	return c.Errors[numberType]
}

func (c *ErrorHandler) Error() string {
	msg := "Errors:\n"
	for nType, errors := range c.Errors {
		msg += fmt.Sprintf("Number Type: %d\n", nType)
		for i, e := range errors {
			msg += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
		}
	}
	return msg
}

func (c *ErrorHandler) FailOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}

func (c *ErrorHandler) HandleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
