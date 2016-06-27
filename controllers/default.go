package controllers

import (
	"github.com/astaxie/beego"
	// "fmt"
)

type MainController struct {
	beego.Controller
}
// @router /all [get]
func (c *MainController) Get() {
	
	//输出json
	c.Data["json"] = map[string]string{"uid": "aaa"}
	c.ServeJSON()
	// fmt.Println("----------------------")
	//渲染模版
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/
	//输出到页面
	// c.Ctx.WriteString("hello")
}
