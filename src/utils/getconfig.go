package utils

import (
	"encoding/xml"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"bytes"
	"github.com/astaxie/beego"
	"net/http"
	"models"
	"encoding/json"
	"fmt"
)

type ConfigElement struct {
	Config1 string
	Config2 string
	Config3 string
	Feature1 string
	Feature2 string
	Feature3 string
}


func GetYamlConfig(path string) map[interface{}]interface{}{
	data, err := ioutil.ReadFile(path)
	m := make(map[interface{}]interface{})
	if err != nil {
		Log(err,beego.LevelWarning)
	}
	err = yaml.Unmarshal([]byte(data), &m)
	return m
}

func GetXMLConfig(path string) map[string]string {
	var t xml.Token
	var err error

	Keylst := make([]string,6)
	Valuelst:=make([]string,6)

	map1:=make(map[string]string)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		Log(err, beego.LevelWarning)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	i:=0
	j:=0
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			Keylst[i]=string(name)
			i=i+1
		case xml.CharData:
			content1 := string([]byte(token))
			Valuelst[j]=content1
			j=j+1
		}
	}
	for count:=0;count<len(Keylst);count++{
		map1[Keylst[count]]=Valuelst[count]
	}

	return map1
}

func GetElement(key string,themap map[interface{}]interface{})string {
	if value,ok:=themap[key];ok {

		return fmt.Sprint(value)
	}

	Log("can't find the config file",beego.LevelWarning)
	return ""
}
//get corpid
func GetCorpId()(corp string,secret string){
	configmsg :=GetYamlConfig("./conf/id_relative.yaml")
	corp =GetElement("corpid",configmsg)
	secret =GetElement("corpsecret",configmsg)
	return corp,secret
}

func GetPubId()(pubid string, pubsecret string){
	configmsg :=GetYamlConfig("./conf/id_relative.yaml")
	pubid =GetElement("pubid",configmsg)
	pubsecret =GetElement("pubsecret",configmsg)
	return pubid, pubsecret

}

func ReadToken()string{
	configmsg :=GetYamlConfig("./conf/id_relative.yaml")
	token := GetElement("token",configmsg)
	return token
}
//get aeskey
func ReadAesKey()string{
	configmsg :=GetYamlConfig("./conf/id_relative.yaml")
	token := GetElement("key",configmsg)
	return token
}


func GetToken()string{
	corp,secret := GetCorpId()
	resp,err :=http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="+corp+"&corpsecret="+secret)
	if err!=nil{
		Log(err,beego.LevelWarning)
	}
	fetchtoken,err :=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var token models.Token
	json.Unmarshal(fetchtoken,&token)
	return token.AccessToken
}