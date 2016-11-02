package controllers

import (
	"github.com/astaxie/beego"
	"models"
	"fmt"
	"encoding/xml"
	"time"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"utils"
)

type PubController struct  {
	beego.Controller
}


//TODO 这里负责回调模式的验证
func (c * PubController) Get() {
	var signature models.Signature
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	fmt.Println(signature.Echostr)
	c.Ctx.WriteString(signature.Echostr)

}




func (c * PubController) Post(){

	var msgIn models.PubTextMsg
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgIn)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}

	if(msgIn.MsgType == "event"){
		msgback := "感谢您的关注(O w O～～下面是功能菜单\n" +
		"1.翻译文字\n"+
		"2.历史上的今天\n"+
		"3.菜谱查询\n"+
		"4.提交报名表联系作者\n"+
		"请输入数字选择使用功能～～输入 菜单 查看菜单"

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

	}else if(msgIn.MsgType == "text"){
		if (msgIn.Content == "1"){
			msgback := "使用翻译功能，请输入 翻译＋翻译内容，如： 翻译 i love you \n"+
			"(可以自动识别输入文字所属语言，并翻译成汉语)"

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

		}else if(strings.HasPrefix(msgIn.Content,"翻译") == true){
			sentence := strings.Replace(msgIn.Content,"翻译","",1)
			appid := "20160909000028452"
			salt := "1435660288"
			key :=  "ze9KluGkjylrnIAAcVl6"
			md5 :=utils.CreateTransString(appid,sentence,salt,key)

			url := `http://api.fanyi.baidu.com/api/trans/vip/translate?q=`+sentence+`&from=auto&to=zh&appid=`+appid+`&salt=`+salt+`&sign=`+ md5
			resp,err :=http.Get(url)
			if err!=nil{
				Lg(err)
			}
			fetchrs,err :=ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			var rs models.TranResult
			json.Unmarshal(fetchrs,&rs)
			Lg(rs)
			var msgback string
			if(rs.ErrorCode == ""){
				msgback = "翻译结果为：" + rs.Translation[0].Dst
			}else {
				msgback = "翻译出错，错误码：" + rs.ErrorCode
			}
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
		}else if(msgIn.Content == "菜单"){
			msgback := "功能菜单如下：\n" +
				"1.翻译文字\n"+
				"2.历史上的今天\n"+
				"3.菜谱查询\n"+
				"4.提交报名表联系作者\n"+
				"请输入数字选择使用功能～～"

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

		}else if(msgIn.Content == "2"){
			msgback := utils.RobotApi("历史上的今天")
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

		} else if(msgIn.Content == "3"){
			msgback := "查询菜谱请输入：查询＋菜名，如 查询 锅包肉"
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

		}else if(strings.HasPrefix(msgIn.Content,"查询") == true){
			sentence := strings.Replace(msgIn.Content,"查询","",1)
			sentence = strings.TrimSpace(sentence)
			msgback := utils.RobotApi(sentence + "的做法")
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

		} else if(msgIn.Content == "4"){
			msgback := `<a href="http://www.nuckylu.com/sample/infome.html">联系开发者</a>`
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

		}else{
			msgback := utils.RobotApi(msgIn.Content)+"\n"+
			"(现在处于聊天模式，输入菜单 查看功能菜单)"
			msgOut := models.PubTextOut{
				ToUserName:msgIn.FromUserName,
				FromUserName:msgIn.ToUserName,
				CreateTime:time.Now().Unix(),
				MsgType:"text",
				Content:fmt.Sprint(msgback),
			}

			xmlData, err := msgOut.ToXml()
			if err != nil {
				c.Abort("500")
			}
			c.Ctx.WriteString(string(xmlData))
		}
	}else {
		msgback := "不能识别您发送的信息类型 （＝ ＝）请重新发送。。。。。 "
		msgOut := models.PubTextOut{
			ToUserName:msgIn.FromUserName,
			FromUserName:msgIn.ToUserName,
			CreateTime:time.Now().Unix(),
			MsgType:"text",
			Content:fmt.Sprint(msgback),
		}

		xmlData, err := msgOut.ToXml()
		if err != nil {
			c.Abort("500")
		}
		c.Ctx.WriteString(string(xmlData))
	}

}

