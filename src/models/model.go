package models
import "encoding/xml"

type Signature struct {
	Msg_signature string `form:"msg_signature"`
	Timestamp string `form:"timestamp"`
	Nonce int `form:"nonce"`
	Echostr string `form:"echostr"`
}


type Enrollinfo struct {
	Name string `form:"name"`
	Sex string `form:"sex"`
	Email string `form:"email"`
	Wechat string `form:"wechat"`

}

type PubTextMsg struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	MsgId int `xml:"MsgId"`
	Content string `xml:"Content"`
	Event string `xml:"Event"`
}


type PubTextOut struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
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

type MsgPlain1 struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
}


type SendDecryptDate struct {
	XMLName xml.Name `xml:"xml"`
	Encrypt string `xml:"Encrypt"`
	MsgSignature string `xml:"MsgSignature"`
	TimeStamp int64 `xml:"TimeStamp"`
	Nonce string `xml:"Nonce"`
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

type Token struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires_in"`
}

type PushMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	AgentID int   	`json:"agentid"`
	Text    TextMsgContent `json:"text"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}

type TranResult struct {
	ErrorCode string `json:"error_code"`
	Translation [1]RsforTrans `json:"trans_result"`
}

type RsforTrans struct{
	Src string `json:"src"`
	Dst string `json:"dst"`

}