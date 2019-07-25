package service

import (
	"net/smtp"
	"os"
	"reflect"

	"github.com/hhhhp52/webtest/src/domain"
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

// CreateAccount is to create a new account
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

	//create user,here set data
	agent = &domain.User{}
	agent.Account = value.FieldByName("Account").String()
	agent.Password = value.FieldByName("Password").String()
	agent.Name = value.FieldByName("Name").String()
	agent.Nickname = value.FieldByName("Nickname").String()
	agent.Email = value.FieldByName("Email").String()

	//here add data to database
	userdao.New(tx, agent)

	tx.Commit()

	Account := os.Getenv("Gmail_Account")
	Password := os.Getenv("Gmail_Password")
	logger.Debug(Account)
	logger.Debug(Password)
	gmailAuth := smtp.PlainAuth("", Account, Password, "smtp.gmail.com")
	to := []string{agent.Email}
	msg := []byte(
		"Subject: This is a test mail!\r\n" +
			"From: test@example.com\r\n" +
			`Content-Type: multipart/mixed; boundary="qwertyuio"` + "\r\n" +
			"\r\n" +
			"--qwertyuio\r\n" +
			"This is the body of email.\r\n" +
			"\r\n" +
			"--qwertyuio--\r\n",
	)

	err := smtp.SendMail("smtp.gmail.com:587", gmailAuth, Account, to, msg)

	if err != nil {
		logger.Debug(err)
	}

	return "success", nil
}
