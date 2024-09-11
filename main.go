package main

import (
	"log"

	"github.com/PyMarcus/mvc-arch/src/configuration/logger"
	"github.com/PyMarcus/mvc-arch/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Fail to get .env", err, zapcore.Field{
			Key: "caller", String: "main",
		})
	}
	logger.Info("Running...")

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(); err != nil {
		logger.Error("Fail to init routes", err, zapcore.Field{
			Key: "caller", String: "main",
		})
		log.Fatal(err)
	}
}
