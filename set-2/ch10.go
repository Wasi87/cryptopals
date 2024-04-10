package set2

import (
	"bytes"
	"crypto/aes"

	set1 "cryptopals/set-1"
	"cryptopals/util"
)

func Chal10(key, filePath string, iv byte) string {
	content := set1.ReadFile(filePath)
	decodedContent := set1.DecodeBase64(content)
	dec := AesDecryptedCBC([]byte(key), decodedContent, iv)
	return string(dec)
}


func AesDecryptedCBC(key, cipherText []byte, iv byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	bs := block.BlockSize()
	undoKey := make([]byte, len(cipherText))
	var undoXor []byte
	
	firstBlock := bytes.Repeat([]byte{iv}, bs)
	for i := 0; i < len(cipherText); i+=bs {
		block.Decrypt(undoKey[i:i+bs],cipherText[i:i+bs])
		undoXor = append(undoXor, set1.Xor(firstBlock, undoKey[i:i+bs])...)
		firstBlock = cipherText[i:i+bs]
	}
	return undoXor
}


func AesEncryptedCBC(key, plainText []byte, iv byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	
	bs := block.BlockSize()
	plainText = util.PKCS7Pad(plainText, bs)
	enc := make([]byte, len(plainText))

	prevBlock := bytes.Repeat([]byte{iv}, bs)
	for i := 0; i < len(plainText); i+=bs {
		xord := set1.Xor(prevBlock, plainText[i:i+bs])
		block.Encrypt(enc[i:i+bs], xord)
		prevBlock = enc[i:i+bs]
	}
	return enc
}