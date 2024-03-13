package set1

import (
	"fmt"
	"log"
	"testing"
)

func TestAes128DecryptECB(t *testing.T) {
	key := "YELLOW SUBMARINE"
	filePath := "texts/ch7.txt"
	result, err := Aes128DecryptECB(key, filePath)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("challenge 7 test:", result)
	// openssl enc -d -aes-128-ecb -base64 -in 7.txt -out ch7_decrypted_file.txt -K 59454C4C4F57205355424D4152494E45
}