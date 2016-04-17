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


func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	block, err := aes.NewCipher(key)
	if err != err {
		return "", err
	}
	blockSize := block.BlockSize()
	fmt.Println("the size is :",blockSize)
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Deallength(input []byte)([]byte,error){
	buf := bytes.NewBuffer(input[16:20])
	var length int32
	err:=binary.Read(buf, binary.BigEndian, &length)
	return input[20 : 20+length],err

}

func NormalDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
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

func AesEncrypt1(plainData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey)
	if len(plainData)%k != 0 {
		plainData = PKCS5Padding(plainData, k)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipherData := make([]byte, len(plainData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, plainData)

	return cipherData, nil
}