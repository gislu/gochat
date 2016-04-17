package models
import "encoding/xml"

type Signature struct {
	Msg_signature string `form:"msg_signature"`
	Timestamp string `form:"timestamp"`
	Nonce int `form:"nonce"`
	Echostr string `form:"echostr"`
}

type MsgPlain struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	AgentID int `xml:"AgentID"`
	Content string `xml:"Content"`
}

type MsgEvent struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Event string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	AgentID int `xml:"AgentID"`
	Content string `xml:"Content"`
}




type MsgCat struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	EventKey string `xml:"EventKey"`
	AgentID int `xml:"AgentID"`
	Content string `xml:"Content"`

}

type DecryptDate struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	Encrypt string `xml:"Encrypt"`

}

type SendDecryptDate struct {
	XMLName xml.Name `xml:"xml"`
	Encrypt string `xml:"Encrypt"`
	MsgSignature string `xml:"MsgSignature"`
	Nonce string `xml:"Nonce"`
	Timestamp string `xml:"TimeStamp"`
}
type Meta struct {
	Category string `json:"category"`
	Catelog string  `json:"catelog"`
}

type Entry struct {
	Meta Meta `json:"meta"`
	Content interface{} `json:"Content"`
}
type Token struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires_in"`
}