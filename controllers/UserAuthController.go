package controllers

import (
	"github.com/astaxie/beego"
)

type UserAuthController struct {
	beego.Controller
}

func (c *UserAuthController) Get() {
	c.TplName = "userAuth/login.tpl"
}
