package core

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// set default database
	// orm.RegisterDataBase("default", "mysql", "root:root@/mybeegodemo?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_pass")+"@/"+beego.AppConfig.String("db_name")+"?charset=utf8", 30)
}
