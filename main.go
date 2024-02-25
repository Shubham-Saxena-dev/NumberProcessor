package main

import (
	"CARIAD/config"
	"CARIAD/internal/cache"
	"CARIAD/pkg/controllers"
	"CARIAD/pkg/service"
	"CARIAD/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	serv       service.NumberService
	controller controllers.NumberController
)

func init() {
	config.InitFromFile(".env")
}

// @title CARIAD
// @version 1.0
// @description Take home test exercise.
// @host localhost:8080
// @schemes http
func main() {
	log.Info("Hi, this is CARIAD take home test")
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()
	routes.RegisterHandlers(server, controller).RegisterHandlers()
	err := server.Run()
	if err != nil {
		gin.Logger()
		log.Error(err)
		os.Exit(1)
	}
}

func initializeLayers() {
	serv = service.NewNumberService(cache.NewCacheService())
	controller = controllers.NewController(serv)
}
