package utils

import (

	"crypto/md5"
	"encoding/hex"
	"net/http"
	"io/ioutil"
	"fmt"
)


func CreateTransString(appid string,query string,salt string,key string,)string{
	material := appid + query + salt + key
	hasher := md5.New()
	hasher.Write([]byte(material))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RobotApi(keymsg string)string{
	url := "http://api.douqq.com/?key=PUVLKzdjeDduTWNHUFVXQUU3PWhRTytOekFrQUFBPT0&msg=" + keymsg
	resp,err :=http.Get(url)
	if err!=nil{
		fmt.Print(err)
	}
	fetchrs,err :=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(fetchrs)

}