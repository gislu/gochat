package controllers

import (
	"github.com/astaxie/beego"
	"models"
	gutils "utils"
	"encoding/xml"
	"time"
	"os"
	"log"
	"fmt"

)




type CorpController struct  {
	beego.Controller
}
//这里填上在企业号回调模式中设置的AESkey




//这里负责回调模式的验证
func (c * CorpController) Get() {
	key :=gutils.ReadAesKey()
	//token := gutils.ReadToken()
	signature := models.Signature{}
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}


	newkey := gutils.Base64Dncode(key)
	rand_msg,err := gutils.AesDecrypt(string(signature.Echostr),newkey)
	if err != nil {
		Lg(err)
	}

	//Lg("After Decrypt we get the result:",string(rand_msg))
	c.Ctx.WriteString(string(rand_msg))
}



func (c * CorpController) Post() {

	key :=gutils.ReadAesKey()
	token := gutils.ReadToken()
	//这里接受微信那边发送来的信息并解析
	var msgDecrypt models.DecryptDate
	var msgIn models.MsgCat
	var sendOut models.SendDecryptDate
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgDecrypt)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}

	//微信那边发来的信息是经过AES加密后的BASE64编码，这里进行解密
	newkey := gutils.Base64Dncode(key)
	rand_msg,err := gutils.AesDecrypt(string(msgDecrypt.Encrypt),newkey)

	if err != nil {
		Lg(err)
	}

	//Lg("Get msg!",string(rand_msg))
	Lg("Get a msg!")
	err1 :=xml.Unmarshal(rand_msg,&msgIn)
	if err1 != nil {
		Lg(err1)
	}

//TODO 通过msgIn.MsgType 判断发送来消息的类型

if msgIn.MsgType=="event"{

	Lg(string(rand_msg))

	msgback := "在这里写好回复给微信的内容,可以自己写函数单门处理业务内容"

	//TODO 这里把回复的消息进行封装
	msgOut := models.CorpEventBackMag{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:fmt.Sprint(msgback),
		AgentID:msgIn.AgentID,
	}


	xmlData ,err := msgOut.ToXml()
	if err != nil {
		c.Abort("500")
	}

	c.Ctx.WriteString(string(xmlData))







}else if msgIn.MsgType=="text"{

	msgback  :="已接单，业务联系人:"+msgIn.FromUserName

	//这里把回复的消息进行封装
	msgOut := models.CorpTextBackMsg{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:msgback,

	}

	xmlData ,err := msgOut.ToXml()
	if err != nil {
		c.Abort("500")
	}

	msg_encrypt,err := gutils.AesEncrypt(string(xmlData),newkey)
	if err != nil {
		Lg(err)
	}

	timeStamp := time.Now().Unix()
	nonce :=fmt.Sprintf("%d", timeStamp)
	sign := gutils.MsgSign(token,nonce,nonce,msg_encrypt)

	sendOut.Encrypt = msg_encrypt
	sendOut.TimeStamp = timeStamp
	sendOut.Nonce =nonce
	sendOut.MsgSignature = sign



	xmlData1 ,err := sendOut.ToXml()
	if err != nil {
		c.Abort("500")
	}


	c.Ctx.WriteString(string(xmlData1))
}


}
func Lg(v ...interface{}) {

	logfile,err:= os.OpenFile("server.log",os.O_RDWR|os.O_APPEND|os.O_CREATE,0);
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return    }
	log.Println(v...)
	logger := log.New(logfile,"\r\n",log.Ldate|log.Ltime);
	logger.SetPrefix("[Info]")
	logger.Println(v...)
	defer logfile.Close();
}

