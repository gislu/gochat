package controllers

import (
	"github.com/astaxie/beego"
	"gochatserver/models"
	gutils "gochatserver/utils"
	"encoding/xml"
	"time"
	"os"
	"log"
	"fmt"

)




type TestController struct  {
	beego.Controller
}
//这里填上在企业号回调模式中设置的AESkey
const 	key = "R44LRQmvEF45rIxY4pItIKQgD4Lh4DBSWVQJPHbzyMM"
const token = "zDzEnkKf65gpADnwJF7yCcROu2"
//TODO 这里负责回调模式的验证
func (c * TestController) Get() {
	signature := models.Signature{}
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	//de64 := gutils.Base64Dncode(signature.Echostr)
	newkey := gutils.Base64Dncode(key)
	//rand_msg,err := gutils.NormalDecrypt(de64,newkey)
	rand_msg,err := gutils.TestDecrypt(string(signature.Echostr),newkey)
	if err != nil {
		Lg(err)
	}

	Lg("After Decrypt we get the result:",string(rand_msg))
	c.Ctx.WriteString(string(rand_msg))
}



func (c * TestController) Post() {


	//TODO: 这里接受微信那边发送来的信息并解析
	var msgDecrypt models.DecryptDate
	var msgIn models.MsgCat
	var sendOut models.SendDecryptDate
	//err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgIn)
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgDecrypt)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}

	//TODO: 微信那边发来的信息是经过AES加密后的BASE64编码，这里进行解密

	//Lg(de64)
	//de64 :=gutils.Base64Dncode(msgDecrypt.Encrypt)
	newkey := gutils.Base64Dncode(key)
	rand_msg,err := gutils.TestDecrypt(string(msgDecrypt.Encrypt),newkey)

	if err != nil {
		Lg(err)
	}

	Lg("Get msg!",string(rand_msg))
	err1 :=xml.Unmarshal(rand_msg,&msgIn)
	if err1 != nil {
		Lg(err1)
	}
	Lg(msgIn.Content,msgIn.MsgType)
	//TODO msgIn.Content 即为微信发来的消息内容

	//TODO 类似的还有msgIn.ToUserName 等


//TODO 通过msgIn.MsgType 判断发送来消息的类型

if msgIn.MsgType=="event"{

	fmt.Println("发来的是事件类信息")

	msgback := "在这里写好回复给微信的内容,可以自己写函数单门处理业务内容"

	//TODO 这里把回复的消息进行封装
	msgOut := models.MsgPlain{
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

	fmt.Println("发来的是文字类信息，发送的内容是:"+msgIn.Content)

	msgback := "在这里写好回复给微信的内容,可以自己写函数单门处理业务内容"

	//TODO 这里把回复的消息进行封装

	msgOut := models.MsgPlain1{
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

	Lg("this is the core xml:"+ string(xmlData))
	msg_encrypt,err := gutils.AesEncrypt(string(xmlData),newkey)
	if err != nil {
		Lg(err)
	}

	timestamp := time.Now().Unix()
	nonce :=fmt.Sprintf("%d",timestamp)
	sign := gutils.MsgSign(token,nonce,nonce,msg_encrypt)

	sendOut.Encrypt = msg_encrypt
	sendOut.TimeStamp = timestamp
	sendOut.Nonce =nonce
	sendOut.MsgSignature = sign



	xmlData1 ,err := sendOut.ToXml()
	if err != nil {
		c.Abort("500")
	}

	Lg("this is what we feedback:\n",string(xmlData1))
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

