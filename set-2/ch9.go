package set2

import (
	"bytes"
)

func PKCS7Pad(text []byte, blockSize int) []byte {
	// 餘 0 仍要填充一個完整的 block
	padLen := blockSize - (len(text) % blockSize)
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	padded := append(text, padding...)
	return padded
}

