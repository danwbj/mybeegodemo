package main

import (
    _ "mybeegodemo/app/modules/core"
	_ "mybeegodemo/app/modules/blog"
	_ "mybeegodemo/app/modules/user"
	_ "mybeegodemo/routers"
	"github.com/astaxie/beego"
)
func main() {
	beego.Run()
}

