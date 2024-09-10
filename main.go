package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PyMarcus/mvc-arch/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(os.Getenv("TEST"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
