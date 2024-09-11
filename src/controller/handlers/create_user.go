package handlers

import (
	"github.com/PyMarcus/mvc-arch/src/configuration/logger"
	"github.com/PyMarcus/mvc-arch/src/configuration/validator"
	"github.com/PyMarcus/mvc-arch/src/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errObj := validator.ValidateUserError(err)
		logger.Error("Fail to create user", err, zapcore.Field{
			Key: "caller", String: "CreateUser",
		})
		c.JSON(errObj.Code, errObj)
	}
}
