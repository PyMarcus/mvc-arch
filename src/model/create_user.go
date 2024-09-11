package model

import (
	"github.com/PyMarcus/mvc-arch/src/configuration/errs"
	"github.com/PyMarcus/mvc-arch/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *errs.Errs {
	logger.Info("Init createuser model", zap.String("caller", "CreateUser"))
	return nil
}
