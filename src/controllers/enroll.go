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
	pushurl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="+gutils.GetToken()

	re1 :=info.City1
	var valid = regexp.MustCompile("[0-9]+")
	city1 :=valid.FindAllStringSubmatch(re1,-1)
	citynum,err:= strconv.Atoi(fmt.Sprint(city1[0][0]))
	if err != nil{
		citynum = 0
	}

	context :="####新用户报名####"+
		"\n用户姓名："+ info.Name +
		"\n用户性别:"+info.Sex +
		"\n用户年龄:"+info.Age+
		"\n目标数量:"+info.Num+
		"\n期待导游爱好:"+info.Love+
		"\n用户目标城市:"+info.City1+
		"\n期望车况:"+info.Car+
		"\n接机时间:"+info.Intime+
		"\n送机时间"+info.Outtime+
		"\n用户手机："+info.Phone+
		"\n用户邮箱或微信:"+info.Email

	Lg(context)

	psMsg := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum,
		Text:  models.TextMsgContent{Content: context},
	}


	body, err := json.MarshalIndent(psMsg, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}



	c.Ctx.WriteString("信息提交成功！")
}


