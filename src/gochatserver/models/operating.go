package models
import (
	"fmt"
	"encoding/xml"
	"encoding/json"
)

func (this *MsgPlain)ToXml() ( []byte , error ) {
	return getXmlData(this)
}

func (this *SendDecryptDate)ToXml() ( []byte , error ) {
	return getXmlData(this)
}

func (this *Entry)ToJson() ( []byte , error ){
	return getJsonData(this)
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
	xmlStr := fmt.Sprintf("%s%s",xml.Header,string(data))
	return []byte(xmlStr) , nil
}
