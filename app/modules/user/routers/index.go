package routers

import (
	"mybeegodemo/app/modules/user/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//初始化namespace
	ns := 
	beego.NewNamespace("/v1",
	beego.NSInclude(
				&controllers.UserController{},
			),
		beego.NSNamespace("/user",
		beego.NSRouter("/users", &controllers.UserController{}, "post:CreateUser"),
		),
	)
	//注册namespace
	beego.AddNamespace(ns)
}
