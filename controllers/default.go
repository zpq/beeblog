package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["isHome"] = "true"
	c.Data["isLogin"] = checkAccout(c.Ctx)
	c.TplName = "home.html"
}
