package controllers

/*
This is a controller class where all routes are directed to.
*/

import (
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/service"
	"CARIAD/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NumberController interface {
	GetNumbersHandler(*gin.Context)
}

type controller struct {
	service service.NumberService
}

func NewController(service service.NumberService) NumberController {
	return &controller{
		service: service,
	}
}

// GetNumbersHandler
// @Summary get numbers slice
// @Description hit the query param api and return merged non-duplicate and sorted result
// @Produce json
// @Param u query []string true "urls" example(localhost:8090/primes)
// @Success 200 {array} int
// @Failure 400 {array} int
// @Router /numbers [get]
func (c *controller) GetNumbersHandler(ctx *gin.Context) {
	errCollector := customerrors.NewErrorCollector()
	urls := ctx.QueryArray("u")
	mergedNumbers := c.service.GetNumbersFromUrl(utils.CreateNumberRequest(urls), errCollector)

	if len(errCollector.Errors) > 0 {
		log.Errorf(fmt.Sprintf("Error occurred: %v", errCollector.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"numbers": mergedNumbers})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"numbers": mergedNumbers})
	}
}
