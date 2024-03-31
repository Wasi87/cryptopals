package set1

import (
	"crypto/aes"
	"cryptopals/util"
	"encoding/base64"
	"log"
	"os"
)


func ReadFile(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil{
		log.Fatal(err)
	}
	return content
}

func DecodeBase64(content []byte) []byte {
	decodedContent, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		log.Fatal("DecodeBase64 error", err)
	}
	return decodedContent
}

func AesDecryptECB(key, cipherText []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()

	decrypted := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i+=blockSize {
		block.Decrypt(decrypted[i:i+blockSize], cipherText[i:i+blockSize])
	}
	return util.PKCS7Unpad(decrypted)
}


func Chal7(key, filePath string) string {
	content := ReadFile(filePath)
	decodedContent := DecodeBase64(content)
	decrypted := AesDecryptECB([]byte(key), decodedContent)

	return string(decrypted)
}