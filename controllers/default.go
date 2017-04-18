package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/bee-test?charset=utf8", 30)
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	var users []orm.Params
	num, err := o.QueryTable("users").Values(&users)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range users {
			fmt.Println(m["Id"], m["Name"])
		}
	}

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["users"] = users
	c.TplName = "index.tpl"
}
