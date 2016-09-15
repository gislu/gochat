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

	keyList := make([]string,6)
	valueList :=make([]string,6)

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
			keyList[i]=string(name)
			i=i+1
		case xml.CharData:
			content1 := string([]byte(token))
			valueList[j]=content1
			j=j+1
		}
	}
	for count:=0;count<len(keyList);count++{
		map1[keyList[count]]= valueList[count]
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

func GetPubId()(pubId string, pubSecret string){
	configmsg :=GetYamlConfig("./conf/id_relative.yaml")
	pubId =GetElement("pubid",configmsg)
	pubSecret =GetElement("pubsecret",configmsg)
	return pubId, pubSecret

}

func ReadToken()string{
	configMsg :=GetYamlConfig("./conf/id_relative.yaml")
	token := GetElement("token", configMsg)
	return token
}
//get aeskey
func ReadAesKey()string{
	configMsg :=GetYamlConfig("./conf/id_relative.yaml")
	token := GetElement("key", configMsg)
	return token
}


func GetCorpToken()string{
	corp,secret := GetCorpId()
	resp,err :=http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="+corp+"&corpsecret="+secret)
	if err!=nil{
		Log(err,beego.LevelWarning)
	}
	fetchToken,err :=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var token models.Token
	json.Unmarshal(fetchToken,&token)
	return token.AccessToken
}

func GetPubToken()string{
	id,secret := GetPubId()
	resp,err :=http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+id+"&secret="+secret)
	if err!=nil{
		Log(err,beego.LevelWarning)
	}
	fetchToken,err :=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var token models.Token
	json.Unmarshal(fetchToken,&token)
	return token.AccessToken
}
