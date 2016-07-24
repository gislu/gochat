package controllers

import (
	"github.com/astaxie/beego"
	"gochatserver/models"
	"fmt"
	"encoding/xml"
	"time"
)

type PubController struct  {
	beego.Controller
}
//这里填上在企业号回调模式中设置的AESkey




//TODO 这里负责回调模式的验证
func (c * 	PubController) Get() {
	var signature models.Signature
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	fmt.Println(signature.Echostr)
	c.Ctx.WriteString(signature.Echostr)

}




func (c * 	PubController) Post(){
	var msgIn models.PubTextMsg
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgIn)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}

	msgback := "这里是自动回复（O w O）"


	msgOut := models.PubTextOut{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:fmt.Sprint(msgback),
	}

	xmlData ,err := msgOut.ToXml()
	if err != nil {
		c.Abort("500")
	}


	c.Ctx.WriteString(string(xmlData))

}
