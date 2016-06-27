package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

// Model Struct
type Blog struct {
    Id   int
    Nickname string `orm:"size(100)"`
    Content string `orm:"size(100)"`
    // Comment  *Comment  `orm:"rel(fk)"`    //设置一对多关系
}
type Comment struct {
    Id   int
    Name string `orm:"size(100)"`
    Content string `orm:"size(100)"`
}



func init() {
	    // register model
    orm.RegisterModel(new(Blog),new(Comment))

    // create table
    orm.RunSyncdb("default", false, true)
	// orm.RegisterModelWithPrefix("bb_", new(Blog),new(User))
    fmt.Println("--------------models init sucess")
}