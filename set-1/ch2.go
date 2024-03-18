package set1

import (
	"encoding/hex"
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
func Xor(str1, str2 []byte) []byte {
	if len(str1) != len(str2) {
		panic("input lengths are different")
	}

	xord := make([]byte, len(str1))

    for i := 0; i < len(str1); i++ {
        xord[i] = str1[i] ^ str2[i]
    }

	return xord
}
