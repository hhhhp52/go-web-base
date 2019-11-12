package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris/v12"
)

//DeleteHandler to delete account
func DeleteHandler(ctx iris.Context) {
	type rule struct {
		Account string `valid:"required"`
	}

	params := &rule{
		Account: ctx.URLParam("Account"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.DeleteAccount(params)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
