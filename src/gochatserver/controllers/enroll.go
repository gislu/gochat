package controllers
import (
	"github.com/astaxie/beego"
	"gochatserver/models"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	gutils "gochatserver/utils"
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

	context := "用户姓名："+ info.Name +"\n用户手机："+info.Phone+"\n用户邮箱："+info.Email
	psMsg := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: 3,
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


