package route

import (
	"fmt"

	"github.com/hhhhp52/webtest/src/handler"
	"github.com/hhhhp52/webtest/src/middleware"
	"github.com/hhhhp52/webtest/src/utils/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

var addr = fmt.Sprintf("%v:%v", config.Get("server.host"), config.Get("server.port"))

// Run maps the routing path and keeps listening for request
func Run() {
	app := iris.New()

	//CORS
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CorsMiddleware)

	// Error Handling
	//app.Use(handler.ErrorHandler)

	// CORS
	//app.RegisterView(iris.HTML("./views", ".html"))

	v1 := app.Party("/v1")
	{
		//創建新帳號
		user := v1.Party("/user")
		{
			//登入
			user.Post("/login", hero.Handler(handler.LoginHandler))
			//創建新帳號
			user.Post("/create", hero.Handler(handler.CreateHandler))
			//登出
			user.Get("/logout", hero.Handler(handler.LogoutHandler))
			//刪除帳號
			user.Post("/delete", hero.Handler(handler.DeleteHandler))
		}

		home := v1.Party("/home")
		{
			home.Get("/showall", hero.Handler(handler.GetAllMemberHandler))
		}
	}
	app.Run(iris.Addr(addr))
	//app.Run(iris.Addr(":80"))
}
