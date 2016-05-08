package utils

import(
	"sort"
	"crypto/sha1"
	"strings"
	"encoding/base64"
	"crypto/aes"
	"io"
	"crypto/cipher"
	"bytes"
	"fmt"
	"encoding/binary"
	"crypto/rand"
	"errors"
	"encoding/hex"
)


func Base64Dncode(src string)string{
	//input:= base64.StdEncoding.EncodeToString([]byte(src))
	rs, _ := base64.StdEncoding.DecodeString(src+"=")
	return string(rs)

}

func Base64Encode(src []byte)string{
	return base64.StdEncoding.EncodeToString(src)
}



func SendMsgSignature(token,timestamp, nonce, msg_encrypt string) string {
	sl := []string{token, timestamp, nonce, msg_encrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func AesEncrypt(text ,key string) (string, error) {
	message := []byte(text)

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(message))); err != nil {
		return "", err
	}

	msgLen := buf.Bytes()

	randBytes := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		return "", err
	}
	id,_ :=GetID()
	messageBytes := bytes.Join([][]byte{randBytes, msgLen, message, []byte(id)}, nil)

	encoded := PKCS7Encode(messageBytes)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	iv := []byte(key)[:16]

	cbc := cipher.NewCBCEncrypter(c, iv)
	cbc.CryptBlocks(encoded, encoded)

	return base64.StdEncoding.EncodeToString(encoded), nil
}


func Deallength(input []byte)([]byte,error){
	buf := bytes.NewBuffer(input[16:20])
	var length int32
	err:=binary.Read(buf, binary.BigEndian, &length)
	return input[20 : 20+length],err

}

func NormalDecrypt(cipherData1 string, aesKey1 string) ([]byte, error) {
	aesKey := []byte(aesKey1)
	cipherData := []byte(cipherData1)
	k := len(aesKey) //PKCS#7
	if len(cipherData) % k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}

func TestDecrypt(text,key string) ([]byte, error) {
	var msgDecrypt []byte

	deciphered, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil,err
	}
	iv := []byte(key)[:16]
	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(deciphered, deciphered)

	decoded := PKCS7Decode(deciphered)

	buf := bytes.NewBuffer(decoded[16:20])

	var msgLen int32
	binary.Read(buf, binary.BigEndian, &msgLen)

	msgDecrypt = decoded[20 : 20+msgLen]
	return msgDecrypt, nil
}

func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(encryptedMsg))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}