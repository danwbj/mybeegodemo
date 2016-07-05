package core

import (
	/*"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver*/

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//  "postgresql/models"

	_ "github.com/lib/pq"
	//  _ "postgresql/routers"
)

func init() {
	// PostgreSQL 配置
	// orm.RegisterDriver("postgres", orm.DR_Postgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user="+beego.AppConfig.String("db_user")+" password="+beego.AppConfig.String("db_pass")+" dbname="+beego.AppConfig.String("db_name")+" host=127.0.0.1 port=5432 sslmode=disable")
	//mysql 配置
	// set default database
	// orm.RegisterDataBase("default", "mysql", "root:root@/mybeegodemo?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_pass")+"@/"+beego.AppConfig.String("db_name")+"?charset=utf8", 30)
}
