package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "nuckylu.com"
	c.Data["Email"] = "ahlxt@foxmail.com"
	c.TplName = "index.tpl"
}
