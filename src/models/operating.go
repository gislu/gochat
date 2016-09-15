package models
import (
	"fmt"
	"encoding/xml"
	"encoding/json"
	"time"
	"github.com/astaxie/beego"
)

func (this *CorpEventBackMag)ToXml() ( []byte , error ) {
	return getXmlData(this)
}

func (this *CorpTextBackMsg)ToXml() ( []byte , error ) {
	return getXmlData(this)
}


func (this *SendDecryptDate)ToXml() ( []byte , error ) {
	return getXmlData(this)
}

func (this *PubTextOut)ToXml() ( []byte , error ) {
	return getXmlData(this)
}

func getJsonData(object interface{}) ( []byte , error ) {
	data,err := json.Marshal(object)
	if err != nil {
		return nil,err
	}
	return data , nil
}

func getXmlData(object interface{})( []byte , error ){
	data , err := xml.Marshal(object)
	if err != nil {
		return nil ,err
	}
	//xmlStr := fmt.Sprintf("%s%s",xml.Header,string(data))
	xmlStr := fmt.Sprintf("%s",string(data))
	return []byte(xmlStr) , nil
}

func PubSendBack(this * beego.Controller,backMsg string,msgIn PubTextMsg)error{
	msgOut := PubTextOut{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:fmt.Sprint(backMsg),
	}

	xmlData ,err := msgOut.ToXml()
	this.Ctx.WriteString(string(xmlData))
	return err
}