package controllers

import (
  "mybeegodemo/app/modules/blog/models"  
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
	// "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
  c.Ctx.WriteString("Get")
}
func (c *BlogController) GetAll() {
  c.Data["json"]="success"
	c.ServeJSON()
  fmt.Println("GetAll")
  // c.Ctx.WriteString("GetAll")
}

func (c *BlogController)  CreateBlog(){
    nickname := c.GetString("nickname")
    content := c.GetString("content")
  	o := orm.NewOrm()
    blog := models.Blog{Nickname: nickname,Content:content}
    // insert
    id, err := o.Insert(&blog)
    if(err!=nil){
      c.Data["json"]="error"
    }
    c.Data["json"]="success"
    c.ServeJSON()

    fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

