package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris"
)

// LoginHandler user login
func LoginHandler(ctx iris.Context) {
	type rule struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
		IP       string `valid:"ip"`
		ErrMsg   string `valid:"-"`
	}

	params := &rule{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
		IP:       ctx.RemoteAddr(), // TODO: get ip from formbody
		ErrMsg:   "",
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		params.ErrMsg = err.Error()
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.Login(params)

	if err != (*error.Error)(nil) {
		params.ErrMsg = err.Error()
		failed(ctx, err)
		return
	}

	ctx.ViewData("message", result)
	ctx.View("login.html")
	return
}
