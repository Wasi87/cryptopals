package set1

import (
	"encoding/hex"
	"fmt"
	"log"
)

func HexToBytes(hexString string) []byte {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// xor - 兩個string長度相等
func Xor(str1, str2 string) (string, error) {
	str1Bytes := HexToBytes(str1)
	str2Bytes := HexToBytes(str2)
	if len(str1) != len(str2) {
		return "", fmt.Errorf("input lengths are different")
	}

	xordBytes := make([]byte, len(str1Bytes))

    for i := 0; i < len(str1Bytes); i++ {
        xordBytes[i] = str1Bytes[i] ^ str2Bytes[i]
    }

	return BytesToHex(xordBytes), nil
} 