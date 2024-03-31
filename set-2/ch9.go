package set2

import "cryptopals/util"

func Chal9(text []byte, blockSize int) []byte {
	return util.PKCS7Pad(text, blockSize)
}

