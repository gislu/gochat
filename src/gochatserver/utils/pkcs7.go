package utils

// PKCS7Decode 方法用于删除解密后明文的补位字符
func PKCS7Decode(text []byte) []byte {
	pad := int(text[len(text)-1])

	if pad < 1 || pad > 32 {
		pad = 0
	}

	return text[:len(text)-pad]
}

// PKCS7Encode 方法用于对需要加密的明文进行填充补位
func PKCS7Encode(text []byte) []byte {
	const BlockSize = 32

	amountToPad := BlockSize - len(text)%BlockSize

	for i := 0; i < amountToPad; i++ {
		text = append(text, byte(amountToPad))
	}

	return text
}
