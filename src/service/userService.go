package service

import (
	"github.com/hhhhp52/webtest/src/error"
	gormdao "github.com/hhhhp52/webtest/src/persistence/gorm"
	userdao "github.com/hhhhp52/webtest/src/persistence/gorm/userDao"
	"github.com/hhhhp52/webtest/src/utils/logger"
)

func loadAccount(account string) (result string, e *error.Error) {
	tx := gormdao.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	user := userdao.GetByAccount(tx, account)
	if user == nil {
		return "", error.DataNotFoundError("Account " + account)
	}
	return "success", nil
}