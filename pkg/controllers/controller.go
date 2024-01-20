package controllers

/*
This is a controller class where all routes are directed to.
*/

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/service"
	"CARIAD/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NumberController interface {
	GetNumbersHandler(*gin.Context)
}

type controller struct {
	service      service.NumberService
	errCollector *customerrors.ErrorHandler
}

func NewController(service service.NumberService, errorHandler *customerrors.ErrorHandler) NumberController {
	return &controller{
		service:      service,
		errCollector: errorHandler,
	}
}

func (c *controller) GetNumbersHandler(ctx *gin.Context) {
	urls := ctx.QueryArray("u")
	mergedNumbers := c.service.GetNumbersFromUrl(utils.CreateNumberRequest(urls))

	if len(c.errCollector.Errors) > 0 {
		c.errCollector.HandleError(ctx, c.errCollector, http.StatusBadRequest)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"numbers": mergedNumbers})
	}
}
