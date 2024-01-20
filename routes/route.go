package routes

/*
This class is a route handler. From here, the requests are directed towards the controller
*/

import (
	"CARIAD/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type Route interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller controllers.NumberController
}

func RegisterHandlers(engine *gin.Engine, controller controllers.NumberController) Route {
	return &route{
		engine:     engine,
		controller: controller,
	}
}

// RegisterHandlers This is a route handler for various requests
func (r *route) RegisterHandlers() {
	r.engine.GET("/numbers", r.controller.GetNumbersHandler)
}
