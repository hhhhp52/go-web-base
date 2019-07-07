package service

import (
	"reflect"

	"github.com/hhhhp52/webtest/src/error"
	gormdao "github.com/hhhhp52/webtest/src/persistence/gorm"
	userdao "github.com/hhhhp52/webtest/src/persistence/gorm/userDao"
	"github.com/hhhhp52/webtest/src/utils/hash"
)

// Login user login
func Login(params interface{}) (result map[string]interface{}, e *error.Error) {
	tx := gormdao.DB()

	value := reflect.ValueOf(params).Elem()
	user := userdao.GetByAccount(tx, value.FieldByName("Account").String())

	if user == nil {
		return nil, error.LoginError()
	}

	if ok := hash.Verify(value.FieldByName("Password").String(), user.Password); !ok {
		return nil, error.LoginError()
	}

	return 
}
