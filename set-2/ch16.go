package set2

import (
	"bytes"
	set1 "cryptopals/set-1"
	"fmt"
	"strings"
	"sync"
)

var (
	once16 sync.Once
	key16 []byte
)

func init() {
	once16.Do(func() {key16 = GenAESKey(16)})
}

func EncryptAndWrap(input string) []byte{
	prefix := "comment1=cooking%20MCs;userdata="
	suffix := ";comment2=%20like%20a%20pound%20of%20bacon"
	r := strings.NewReplacer(
		"=", "%3D",
		";", "%3B",
	)
	replaced := r.Replace(input)
	wrapped := fmt.Sprintf("%s%s%s", prefix, replaced, suffix)
	encrypted := AesEncryptedCBC(key16, []byte(wrapped), byte(0))
	return encrypted
}

func CheckAdmin(input []byte) bool {
	decrypted := AesDecryptedCBC(key16, input, byte(0))
	if strings.Contains(string(decrypted), "admin") {
		fmt.Println("check admin:", string(decrypted))
		return true
	}
	fmt.Println("check admin:", string(decrypted))	 
	return false 
}

func AddAdmin() []byte {
	insertedContent := "aaaaa;admin=true"
	a := bytes.Repeat([]byte("a"), 32)

	cipherText := EncryptAndWrap(string(a))
	flipper := set1.Xor(cipherText[32:48], []byte(insertedContent))
	flipper = set1.Xor(bytes.Repeat([]byte("a"), 16), flipper)

	padded := append(cipherText[:32], flipper...)
	padded = append(padded, cipherText[48:]...)
	
	return padded
}