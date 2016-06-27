package main

import (
    _ "mybeegodemo/app/modules/core"
	_ "mybeegodemo/app/modules/blog"
	_ "mybeegodemo/app/modules/user"
	_ "mybeegodemo/routers"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql" // import your used driver
	// "fmt"
)

/*// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}*/
/*func init() {
    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}*/

func main() {
/*	o := orm.NewOrm()

    user := User{Name: "dandan"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    u := User{Id: user.Id}
	err = o.Read(&u)*/
    

    
	beego.Run()
}

