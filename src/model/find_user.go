package model

import (
	"github.com/PyMarcus/mvc-arch/src/configuration/errs"
	"github.com/PyMarcus/mvc-arch/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) FindUser(string) (*userDomain, *errs.Errs) {
	logger.Info("Init finduser model", zap.String("caller", "FindUser"))
	return nil, nil
}
