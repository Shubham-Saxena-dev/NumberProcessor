package main

import (
	"CARIAD/config"
	"CARIAD/internal/customerrors"
	"CARIAD/pkg/controllers"
	"CARIAD/pkg/service"
	"CARIAD/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	serv         service.NumberService
	controller   controllers.NumberController
	errCollector *customerrors.ErrorHandler
)

func init() {
	config.InitFromFile(".env")
	errCollector = customerrors.NewErrorHandler()
}

func main() {
	log.Info("Hi, this is CARIAD take home test")
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()
	routes.RegisterHandlers(server, controller).RegisterHandlers()
	err := server.Run("localhost:" + config.EnvConfigs.App.AppPort)
	if err != nil {
		errCollector.FailOnError(err, "Server initialization failed")
	}
}

func initializeLayers() { // we can also have dependency injector
	serv = service.NewNumberService(errCollector)
	controller = controllers.NewController(serv, errCollector)
}
