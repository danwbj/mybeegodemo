package core
import (
    "github.com/astaxie/beego"
    "fmt"
    "github.com/astaxie/beego/context" 
)


var FilterLogin = func(c *context.Context) {
  /*  c.Data["json"]="error"
	c.ServeJSON()*/
    fmt.Println("filters")
    // ctx.Redirect(302, "/login")
    return
}
func init() {
    beego.InsertFilter("/*",beego.BeforeExec,FilterLogin)
}
