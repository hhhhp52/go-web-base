package handler

import (
	"github.com/kataras/iris"
)

// HelloHandler return hello message
func HelloHandler(ctx iris.Context) {
	ctx.ViewData("message", "Hello world!")
	ctx.View("hello.html")
	return
}
