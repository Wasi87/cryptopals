package set2

import (
	"bytes"
	"cryptopals/util"
	"fmt"
	"math/rand"
	"time"
)

func GenAESKey(size int) []byte {
	key := make([]byte, size)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	r.Read(key)
	return key
}


func EncryptionOracle(plainText []byte) ([]byte, int) {
	key := GenAESKey(16)
	plainText = append(plainText, byte(rand.Intn(6)+5))
	plainText = append([]byte{byte(rand.Intn(6)+5)}, plainText...)

	mode := rand.Intn(2)
	var encrypted []byte 
	switch mode {
	case 0:
		fmt.Println("ECB mode")
		encrypted = util.AesEncryptECB(key, plainText)
	case 1:
		fmt.Println("CBC mode")
		encrypted = AesEncryptedCBC(key, plainText, []byte("\x00"))
	}
	return encrypted, mode
}


func Chal11(){
	encrypted, _ := EncryptionOracle(bytes.Repeat([]byte("A"), 48))

	maxCount := util.DetectECB(encrypted, 16)
	if maxCount != 0 {
		fmt.Println("guess ECB")
	} else {
		fmt.Println("guess CBC")
	}
}