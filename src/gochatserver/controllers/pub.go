package controllers

import (
	"github.com/astaxie/beego"
	"gochatserver/models"
	gutils "gochatserver/utils"
	"sort"
	"crypto/sha1"
	"encoding/hex"
)

type PubController struct  {
	beego.Controller
}
const token = "test"
func (c * PubController) Get() {
	signature := models.Signature1{}
	if err := c.ParseForm(&signature) ; err != nil{
		gutils.Log(err,beego.LevelNotice)
		c.Abort("400")
	}
	strs := sort.StringSlice{token, signature.Timestamp, signature.Nonce}
	strs.Sort()
	buf := make([]byte,0)
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	hashsum := sha1.Sum(buf)
	Lg(hex.EncodeToString(hashsum[:]))
	c.Ctx.WriteString(signature.Echostr)
}

func (c * PubController) Post() {

}