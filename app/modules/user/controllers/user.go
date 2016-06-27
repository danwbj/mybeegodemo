package controllers

import (
  "mybeegodemo/app/modules/user/models"  
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
	// "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController)  CreateUser(){
    username := c.GetString("username")
    age := c.GetString("age")
  	o := orm.NewOrm()
    user := models.User{Username: username,Age:age}
    // insert
    id, err := o.Insert(&user)
    if(err!=nil){
      c.Data["json"]="error"
    }
    c.Data["json"]="success"
    c.ServeJSON()

    fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

