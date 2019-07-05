package main

import (
	"github.com/hhhhp52/webtest/src/route"
)

func main() {
	/*
		app := iris.New()
		app.RegisterView(iris.HTML("./views", ".html"))

		app.Get("/", func(ctx iris.Context) {
			ctx.ViewData("message", "Hello world!")
			ctx.View("hello.html")
		})

		app.Get("/v1/login", func(ctx iris.Context) {
			ctx.ViewData("message", "Login!")
			ctx.View("login.html")
		})

		app.Run(iris.Addr(":80"))
	*/
	route.Run()
}
