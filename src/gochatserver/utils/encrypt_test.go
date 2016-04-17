package utils

import (
	"testing"
	"reflect"
)
	const (

		key = "R44LRQmvEF45rIxY4pItIKQgD4Lh4DBSWVQJPHbzyMM"
	)

func Test_Base64(t *testing.T){
		desrc :=Base64Dncode(key)
		t.Log("the decoded param is :",desrc)


		ensrc :=Base64Encode([]byte(desrc))
		t.Log("the encoded para is:" ,ensrc)

	if ensrc !=(key+"="){
		t.Error("the method of en/decode realized in wrong way")
	}
	}

func Test_Aes(t*testing.T){
	testmaterial :="this is a test balabalabala"
		rs,err:=	AesEncrypt(testmaterial,[]byte(Base64Dncode(key)))
		if err != nil{
			t.Error(err)
		}
		rs1 ,err :=AesDecrypt(rs,[]byte(Base64Dncode(key)))

		t.Log(string(rs1))

		t.Log(reflect.DeepEqual(rs1,[]byte(testmaterial)))


}



