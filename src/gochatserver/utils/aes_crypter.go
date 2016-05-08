package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// MessageCrypter 封装了生成签名和消息加解密的方法
type MessageCrypter struct {
	token string
	appID string
	key   []byte
	iv    []byte
}

// NewMessageCrypter 方法用于创建 MessageCrypter 实例
//
// token 为开发者在微信开放平台上设置的 Token，
// encodingAESKey 为开发者在微信开放平台上设置的 EncodingAESKey，
// appID 为企业号的 CorpId 或者 AppId
func NewMessageCrypter(token, encodingAESKey, appID string) (MessageCrypter, error) {
	var key []byte
	var err error

	if key, err = base64.StdEncoding.DecodeString(encodingAESKey + "="); err != nil {
		return MessageCrypter{}, err
	}

	if len(key) != 32 {
		return MessageCrypter{}, errors.New("encodingAESKey invalid")
	}

	iv := key[:16]

	return MessageCrypter{
		token,
		appID,
		key,
		iv,
	}, nil
}

// GetSignature 方法用于返回签名
func (w MessageCrypter) GetSignature(timestamp, nonce, msgEncrypt string) string {
	sl := []string{w.token, timestamp, nonce, msgEncrypt}
	sort.Strings(sl)

	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}

// Decrypt 方法用于对密文进行解密
//
// 返回解密后的消息，CropId/AppId, 或者错误信息
func (w MessageCrypter) Decrypt(text string) ([]byte, string, error) {
	var msgDecrypt []byte
	var id string

	deciphered, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, "", err
	}

	c, err := aes.NewCipher(w.key)
	if err != nil {
		return nil, "", err
	}

	cbc := cipher.NewCBCDecrypter(c, w.iv)
	cbc.CryptBlocks(deciphered, deciphered)

	decoded := PKCS7Decode(deciphered)

	buf := bytes.NewBuffer(decoded[16:20])

	var msgLen int32
	binary.Read(buf, binary.BigEndian, &msgLen)

	msgDecrypt = decoded[20 : 20+msgLen]
	id = string(decoded[20+msgLen:])

	return msgDecrypt, id, nil
}

// Encrypt 方法用于对明文进行加密
func (w MessageCrypter) Encrypt(text string) (string, error) {
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

	messageBytes := bytes.Join([][]byte{randBytes, msgLen, message, []byte(w.appID)}, nil)

	encoded := PKCS7Encode(messageBytes)

	c, err := aes.NewCipher(w.key)
	if err != nil {
		return "", err
	}

	cbc := cipher.NewCBCEncrypter(c, w.iv)
	cbc.CryptBlocks(encoded, encoded)

	return base64.StdEncoding.EncodeToString(encoded), nil
}

