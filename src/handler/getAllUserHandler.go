package handler

import (
	"github.com/kataras/iris"
)

// GetAllUserHandler return hello message
func GetAllUserHandler(ctx iris.Context) {

	ctx.ViewData("message", "Hello world!")
	ctx.View("hello.html")
	return
}
