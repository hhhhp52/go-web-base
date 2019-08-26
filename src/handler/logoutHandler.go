package handler

import (
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris"
)

// LogoutHandler user logout
func LogoutHandler(ctx iris.Context) {

	result, err := service.Logout()

	if err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, result)
	return
}
