package handlers

import (
	"github.com/PyMarcus/mvc-arch/src/configuration/validator"
	"github.com/PyMarcus/mvc-arch/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errObj := validator.ValidateUserError(err)
		c.JSON(errObj.Code, errObj)
	}
}
