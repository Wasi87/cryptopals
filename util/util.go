package util

import (
	"bytes"
	"crypto/aes"
	"log"
)

func PKCS7Pad(text []byte, blockSize int) []byte {
	// 餘 0 仍要填充一個完整的 block
	padLen := blockSize - (len(text) % blockSize)
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	padded := append(text, padding...)
	return padded
}


func PKCS7Unpad(text []byte) []byte {
	padLen := int(text[len(text)-1])

	for i := len(text)-padLen; i < len(text); i++ {
		if text[i] != byte(padLen) {
			log.Fatal("unpadPKCS7: invalid padding")
		}
	}
	return text[:len(text)-padLen]
}


func AesEncryptECB(key, plainText []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	plainText = PKCS7Pad(plainText, bs)
	
	encrypted := make([]byte, len(plainText))
	for i := 0; i < len(plainText); i+=bs {
		block.Encrypt(encrypted[i:i+bs], plainText[i:i+bs])
	}
	return encrypted
}


// bytes to chunks
func BytesToChunks(line []byte, blockSize int) [][]byte {
	chunks := make([][]byte, len(line)/blockSize)
	for i := 0; i < len(line); i += blockSize {
		end := i + blockSize
		if end > len(line) {
			end = len(line)
		}
		chunks[i/blockSize] = line[i:end]
	}
	return chunks
}


func DetectECB(fileContent[]byte, blockSize int) int {
	maxCount := 0

	chunks := BytesToChunks([]byte(fileContent), blockSize)
	for i := 0 ; i < len(chunks)-1 ; i++ {
		count := 0
		for j := i+1 ; j < len(chunks); j++ {
			if bytes.Equal(chunks[i], chunks[j]) {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
