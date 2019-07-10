package handler

import (
	"github.com/kataras/iris"
)

func GetCreateHandler(ctx iris.Context) {
	ctx.View("createaccount.html")
	return
}
