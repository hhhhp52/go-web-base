package route

import (
	"github.com/hhhhp52/webtest/src/handler"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

//var addr = fmt.Sprintf("%v:%v", config.Get("server.host"), config.Get("server.port"))

// Run maps the routing path and keeps listening for request
func Run() {
	app := iris.New()

	// Error Handling
	//app.Use(handler.ErrorHandler)

	// CORS
	//app.AllowMethods(iris.MethodOptions)
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", hero.Handler(handler.HelloHandler))
	{

	}

	v1 := app.Party("/v1")
	{
		//登入
		v1.Get("/login", hero.Handler(handler.LoginHandler))
		{

		}
	}
	//app.Run(iris.Addr(addr))
	app.Run(iris.Addr(":80"))
}