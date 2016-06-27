package routers

import (
	"mybeegodemo/controllers"
    // "mybeegodemo/app/modules/blog/controllers"
    // _ "mybeegodemo/app/modules/blog"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Include(&controllers.BlogController{})
    beego.Router("/", &controllers.MainController{})
    // beego.Router("/sayHello", &controllers.MainController{})
}
