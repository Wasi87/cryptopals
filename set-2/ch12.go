package set2

import (
	"bytes"
	set1 "cryptopals/set-1"
	"cryptopals/util"
	"fmt"
	"slices"
)

// 💀

func Oracle(plainText, key []byte)[]byte {
	text := set1.ReadFile("texts/ch12.txt")
	unknown := set1.DecodeBase64(text) // 138

	plainText = append(plainText, unknown...)
	encrypted := util.AesEncryptECB(key, plainText)
	return encrypted
}

func Chal12()[]byte{

	key := GenAESKey(16)

	// find blockSize
	preLength := 0
	var bs int
	for i := 0;; i++ {
		input := bytes.Repeat([]byte("A"), i)
		enc := Oracle(input, key)
		if len(enc) > preLength && preLength != 0 {
			fmt.Println("block size:", len(enc), preLength)
			bs = len(enc) - preLength
			break
		}
		preLength = len(enc)
	}

	// check ECB mode
	count := util.DetectECB(bytes.Repeat([]byte("A"), 144), bs)
	if count == 0 {
		panic("Not ECB mode")
	}

	// 要拿少一個byte+遍歷 跟 少一個比較
	// 起始為15個 
	// 起始為14+answer ... 到0 

	// 換下一排
	// 從15個A 偵測第二行 ＝ 15個A + 16已知 遍歷最後一位數
	var bytesAns []byte
	for k := 0; k < len(Oracle([]byte{}, key)); k+=bs {
		for i := 15; i >= 0; i-- {
			// 15~0
			// 傳入15 - 15+0+1(最後一個遍歷)
			// 傳入14 - 14+1+1(最後一個遍歷)
			// 傳入0 - 0+15+1
			
			input := bytes.Repeat([]byte("A"), i)
			enc := Oracle(input , key)
			
			for j := 0; j < 255; j++ {
				var inputGuess []byte
				if i > 0 {
					inputGuess = bytes.Repeat([]byte("A"), i)
					} 
					if len(bytesAns) != 0 {
						inputGuess = append(inputGuess, bytesAns...)
					}
					inputGuess = append(inputGuess, byte(j))			
					encGuess := Oracle(inputGuess, key)
					
					if slices.Equal(enc[k:k+bs], encGuess[k:k+bs]) {
						bytesAns = append(bytesAns, byte(j))
						break
					}
				}
			}
		}

	return bytesAns
}
