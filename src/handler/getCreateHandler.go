package handler

import (
	"github.com/kataras/iris"
)

//GetCreateHandler go to create account page
func GetCreateHandler(ctx iris.Context) {
	ctx.View("createaccount.html")
	return
}
