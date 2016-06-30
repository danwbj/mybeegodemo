package routers

import (
	"mybeegodemo/app/modules/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Router("/v1/blog", &BlogController{},"get:GetAll")
	//初始化namespace
	ns := 
	beego.NewNamespace("/v1",
	beego.NSInclude(
				&controllers.BlogController{},
			),
	// beego.NSRouter("/blog", &controllers.BlogController{}, "get:GetAll"),
		beego.NSNamespace("/blog",
		beego.NSRouter("/blogs", &controllers.BlogController{}, "post:CreateBlog"),
		beego.NSRouter("/blogs/:id", &controllers.BlogController{}, "get:ReadBlog"),
		beego.NSRouter("/all", &controllers.BlogController{}, "get:GetAll"),
	/*		beego.NSInclude(
				&controllers.BlogController{},
			),*/
		),
	)
	//注册namespace
	beego.AddNamespace(ns)
}
