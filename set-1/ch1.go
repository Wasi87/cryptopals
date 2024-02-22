package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64(hexString string) string {
	b, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	b64 := base64.RawStdEncoding.EncodeToString(b)
	return b64
}