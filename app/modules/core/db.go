package core

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql" // import your used driver
    "fmt"
)
func init() {
    // set default database
    // orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
    fmt.Println(beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_pass")+"@/"+beego.AppConfig.String("db_name")+"?charset=utf8")
    orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_pass")+"@/"+beego.AppConfig.String("db_name")+"?charset=utf8", 30)
}
