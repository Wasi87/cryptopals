package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)


func genKey(key string) []byte{
	return []byte(key)
}

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

func AesDecryptECB(key, cipherText []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)

	blockSize := block.BlockSize()
	if len(cipherText)%blockSize != 0 {
		return nil, fmt.Errorf("ciphertext length is not a multiple of block size")
	}
	
	decrypted := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i+=blockSize {
		block.Decrypt(decrypted[i:i+blockSize], cipherText[i:i+blockSize])
	}
	return decrypted, nil
}

func Aes128DecryptECB(key, filePath string)(string, error){
	keyBytes := genKey(key)
	if len(keyBytes) != 16 {
		log.Fatal("the length of key is not equal to 16")
	}
	content := ReadFile(filePath)
	decodedContent := DecodeBase64(content)
	decrypted, err := AesDecryptECB(keyBytes, decodedContent)
	if err != nil{
		return "", err
	}

	return string(decrypted), nil
}