package handlers

import (
	"fmt"

	"github.com/PyMarcus/mvc-arch/src/configuration/errs"
	"github.com/PyMarcus/mvc-arch/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errObj := errs.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error: %s!", err.Error()))
		c.JSON(errObj.Code, errObj)
	}
}
