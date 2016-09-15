package controllers
import (
	"github.com/astaxie/beego"
	"models"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"regexp"
	gutils "utils"
	"strconv"
)

type EnrollController struct  {
	beego.Controller
}
func (c * EnrollController) Get(){

	var info models.Enrollinfo
	if err := c.ParseForm(&info) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}


	context :="####新用户报名####"+
		"\n用户姓名："+ info.Name +
		"\n用户性别:"+info.Sex +
		"\n用户年龄:"+info.Wechat+
		"\n用户邮箱或微信:"+info.Email
	Lg(context)
	c.Redirect("/static/QR.html", 301)
}


