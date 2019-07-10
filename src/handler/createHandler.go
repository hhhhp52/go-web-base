package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris"
)

// LoginHandler user login
func CreateHandler(ctx iris.Context) {
	type rule struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
		Name     string `valid:"required"`
		Nickname string `valid:"required"`
	}

	params := &rule{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
		Name:     ctx.FormValue("Name"),
		Nickname: ctx.FormValue("Nickname"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.CreateAccount(params)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	ctx.ViewData("message", result)
	ctx.View("create-redirect.html")
	return
}