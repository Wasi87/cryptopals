package set2

import (
	"bytes"
	"crypto/aes"

	set1 "cryptopals/set-1"
)

func Chal10(key, filePath, iv string) string {
	content := set1.ReadFile(filePath)
	decodedContent := set1.DecodeBase64(content)
	ivByte := bytes.Repeat([]byte(iv), len(key))
	dec := AESDecryptedCBC([]byte(key), decodedContent, ivByte)
	return string(dec)
}


func AESDecryptedCBC(key, cipherText, iv []byte) []byte {
	block, _ := aes.NewCipher(key)

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