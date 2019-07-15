package handler

import (
	"github.com/kataras/iris"
)

// HomeHandler return welcome message
func HomeHandler(ctx iris.Context) {
	ctx.ViewData("message", "歡迎回到主頁")
	ctx.View("home.html")
	return
}
