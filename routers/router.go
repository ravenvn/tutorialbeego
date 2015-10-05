package routers

import (
	"tutorial/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Router("users", &controllers.UserController{}, "get:List")
	beego.Router("user/:id", &controllers.UserController{}, "get:Show")
	beego.Router("create", &controllers.UserController{}, "get:Create")
	beego.Router("save", &controllers.UserController{}, "post:Save")
}
