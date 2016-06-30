package models
import (
	"github.com/astaxie/beego/orm"
    "fmt"
)

// Model Struct
type User struct {
    Id   int
    Username string `orm:"size(100)"`
    Age string `orm:"size(100)"`
}


func init() {
	// register model
    orm.RegisterModel(new(User))
    // create table
    orm.RunSyncdb("default", false, false)
    fmt.Println("--------------user models init success")
    orm.ResetModelCache()
}