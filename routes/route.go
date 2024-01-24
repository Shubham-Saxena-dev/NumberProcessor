package routes

/*
This class is a route handler. From here, the requests are directed towards the controller
*/

import (
	_ "CARIAD/docs"
	"CARIAD/pkg/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.engine.GET("/numbers", r.controller.GetNumbersHandler)
}
