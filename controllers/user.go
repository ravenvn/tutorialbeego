package controllers

import (
	"strconv"
	"tutorial/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) List() {
	o := orm.NewOrm()

	// o.Using("default")
	// user := models.User{id: 1}
	// err := o.Read(&user)

	var users []*models.User
	o.QueryTable("user").All(&users)
	c.Data["users"] = users

	c.TplNames = "users.tpl"
}

func (c *UserController) Show() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	var user models.User
	o.QueryTable("user").Filter("id", id).One(&user)
	c.Data["user"] = user

	c.TplNames = "user.tpl"
}

func (c *UserController) Create() {
	c.TplNames = "create.tpl"
}

func (c *UserController) Save() {
	name := c.GetString("name")
	o := orm.NewOrm()
	var user models.User
	user.Name = name
	o.Insert(&user)
	c.Redirect("/users", 302)
}
