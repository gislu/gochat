package test
import (
	"testing"
	"gochatserver/models"
	"fmt"
	"time"
	"encoding/xml"
	"gochatserver/utils"
	"github.com/astaxie/beego"
)

func TestSerilizeMsgoutJson(t *testing.T) {
//	var meta1 models.Meta
//	meta1.Category = "message"
//	meta1.Catelog = "echo"
//
//
//	b , err := json.Marshal(msgout)
//	if err != nil {
//		t.Error("Seriaze Failed")
//	}
//	t.Log(string(b),"\n")
}

func TestSerilizeMsgoutXml(t *testing.T) {
	var msg2wechat models.MsgPlain
	msg2wechat.Content = fmt.Sprint("测试数据")
	msg2wechat.CreateTime = time.Now().Unix()
	msg2wechat.FromUserName = "呵呵"
	msg2wechat.ToUserName = "哈哈"
	msg2wechat.MsgType = "text"
	data , err := xml.Marshal(msg2wechat)
	if err != nil {
		utils.Log("error",beego.LevelError)
	}
	t.Log(fmt.Sprintf("%s%s",xml.Header,string(data)))
}
