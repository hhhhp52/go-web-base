package service

import (
	"reflect"
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/domain"
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

func CreateAccount(params interface{}) (result string, e *error.Error) {
	tx := gormdao.DB().Begin()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	value := reflect.ValueOf(params).Elem()

	agent := userdao.GetByAccount(tx, value.FieldByName("Account").String())

	if agent != nil {
		return "", error.InvaildValueError("Account Already Exists.")
	}

	//create user
	//here set data
	agent = &domain.User{}
	agent.Account = value.FieldByName("Account").String()
	agent.Password = value.FieldByName("Password").String()
	agent.Name = value.FieldByName("Name").String()
	agent.Nickname = value.FieldByName("Nickname").String()
	
	//here add data to database 
	userdao.New(tx, agent)

	tx.Commit()

	return "success", nil
}

