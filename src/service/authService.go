package service

import (
	"reflect"

	"github.com/hhhhp52/webtest/src/error"
	gormdao "github.com/hhhhp52/webtest/src/persistence/gorm"
	userdao "github.com/hhhhp52/webtest/src/persistence/gorm/userDao"
	//"github.com/hhhhp52/webtest/src/utils/hash"
	"github.com/hhhhp52/webtest/src/utils/logger"
)

// Login user login
func Login(params interface{}) (result []string, e *error.Error) {
	tx := gormdao.DB()

	defer func(){
		if r := recover(); r != nil{
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	value := reflect.ValueOf(params).Elem()
	user := userdao.GetByAccount(tx, value.FieldByName("Account").String())
	
	if user == nil {
		return nil, error.LoginError()
	}
	// 因加密需解密
	/*
	if ok := hash.Verify(value.FieldByName("Password").String(), user.Password); !ok {
		fmt.Println("password : ", user.Password)
		fmt.Println("error here ok")
		return nil, error.LoginError()
	}
	*/
	
	// 密碼比對該方式有錯誤，之後再進行確認 
	if user.Password != value.FieldByName("Password").String(){
		return nil, error.LoginError()
	}
	
	loadAccount(user.Account)

	result = append(result, user.Account)

	return result , nil
}
