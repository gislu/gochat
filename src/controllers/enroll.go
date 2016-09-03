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
	citynum1,err:= strconv.Atoi(fmt.Sprint(city1[0][0]))
	if err != nil{
		citynum1 = 0
	}

	re2 :=info.City2
	city2 :=valid.FindAllStringSubmatch(re2,-1)
	citynum2,err:= strconv.Atoi(fmt.Sprint(city2[0][0]))
	if err != nil{
		citynum2 = 0
	}

	re3 :=info.City3
	city3 :=valid.FindAllStringSubmatch(re3,-1)
	citynum3,err:= strconv.Atoi(fmt.Sprint(city3[0][0]))
	if err != nil{
		citynum3 = 0
	}

	re4 :=info.City4
	city4 :=valid.FindAllStringSubmatch(re4,-1)
	citynum4,err:= strconv.Atoi(fmt.Sprint(city4[0][0]))
	if err != nil{
		citynum4 = 0
	}

	re5 :=info.City5
	city5 :=valid.FindAllStringSubmatch(re5,-1)
	citynum5,err:= strconv.Atoi(fmt.Sprint(city5[0][0]))
	if err != nil{
		citynum5 = 0
	}

	context :="####新用户报名####"+
		"\n用户姓名："+ info.Name +
		"\n用户性别:"+info.Sex +
		"\n用户年龄:"+info.Age+
		"\n目标数量:"+info.Num+
		"\n期待导游爱好:"+info.Love+
		"\n用户目标城市1:"+info.City1+
		"\n用户目标城市2:"+info.City2+
		"\n用户目标城市3:"+info.City3+
		"\n用户目标城市4:"+info.City4+
		"\n用户目标城市5:"+info.City5+
		"\n期望车况:"+info.Car+
		"\n接机时间:"+info.Intime+
		"\n送机时间"+info.Outtime+
		"\n用户手机："+info.Phone+
		"\n用户邮箱或微信:"+info.Email

	Lg(context)

	psMsg1 := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum1+3,
		Text:  models.TextMsgContent{Content: context},
	}

	psMsg2 := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum2+3,
		Text:  models.TextMsgContent{Content: context},
	}

	psMsg3 := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum3+3,
		Text:  models.TextMsgContent{Content: context},
	}

	psMsg4 := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum4+3,
		Text:  models.TextMsgContent{Content: context},
	}

	psMsg5 := &models.PushMsg{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: citynum5+3,
		Text:  models.TextMsgContent{Content: context},
	}

	body, err := json.MarshalIndent(psMsg1, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}


	body2, err := json.MarshalIndent(psMsg2, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body2))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}

	body3, err := json.MarshalIndent(psMsg3, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body3))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}

	body4, err := json.MarshalIndent(psMsg4, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body4))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}

	body5, err := json.MarshalIndent(psMsg5, " ", "  ")
	if err != nil {
		Lg(err)
	}
	_,err =http.Post(pushurl,"application/json",bytes.NewReader(body5))
	if err != nil {
		Lg(err)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
	}



	c.Ctx.WriteString("信息提交成功！")
}


