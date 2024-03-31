package set2

import (
	"bytes"
	"crypto/aes"

	set1 "cryptopals/set-1"
	"cryptopals/util"
)

func Chal10(key, filePath, iv string) string {
	content := set1.ReadFile(filePath)
	decodedContent := set1.DecodeBase64(content)
	ivByte := bytes.Repeat([]byte(iv), len(key))
	dec := AesDecryptedCBC([]byte(key), decodedContent, ivByte)
	return string(dec)
}


func AesDecryptedCBC(key, cipherText, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	bs := block.BlockSize()
	undoKey := make([]byte, len(cipherText))
	var undoXor []byte
	
	firstBlock := iv
	for i := 0; i < len(cipherText); i+=bs {
		block.Decrypt(undoKey[i:i+bs],cipherText[i:i+bs])
		undoXor = append(undoXor, set1.Xor(firstBlock, undoKey[i:i+bs])...)
		firstBlock = cipherText[i:i+bs]
	}
	return undoXor
}


func AesEncryptedCBC(key, plainText, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	
	bs := block.BlockSize()
	plainText = util.PKCS7Pad(plainText, bs)
	enc := make([]byte, len(plainText))

	ivByte := bytes.Repeat([]byte(iv), bs)
	firstBlock := ivByte
	
	for i := 0; i < len(plainText); i+=bs {
		xord := set1.Xor(firstBlock, plainText[i:i+bs])
		block.Encrypt(enc[i:i+bs], xord)
		firstBlock = enc[i:i+bs]
	}
	return enc
}