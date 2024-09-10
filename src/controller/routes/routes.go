package routes

import (
	"github.com/PyMarcus/mvc-arch/src/controller/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/getUserById/:userId", handlers.GetUserById)
	rg.GET("/getUserByEmail/:userEmail", handlers.GetUserByEmail)
	rg.POST("/createUser", handlers.CreateUser)
	rg.PUT("/updateUser", handlers.UpdateUser)
	rg.DELETE("/deleteUser", handlers.DeleteUser)
}
