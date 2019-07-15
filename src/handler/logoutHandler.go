package handler

import (
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris"
)

// LogoutHandler user logout
func LogoutHandler(ctx iris.Context) {

	_, err := service.Logout()

	if err != nil {
		failed(ctx, err)
		return
	}

	ctx.View("logout.html")
	return
}
