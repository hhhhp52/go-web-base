package handler

import (
	"github.com/kataras/iris"
)

// LoginHandler user login
func LoginHandler(ctx iris.Context) {
	ctx.ViewData("message", "Login!")
	ctx.View("login.html")
	return
}
