package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "goeverywhere"
	c.Data["Email"] = "ahlxt@foxmail.com"
	c.TplName = "提交信息.html"
}