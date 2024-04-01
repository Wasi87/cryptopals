package set2

import (
	"bytes"
	"fmt"
)


func VerifyPadding(text []byte, blockSize int) ([]byte, error) {
	if len(text)%blockSize != 0 {
		return nil, fmt.Errorf("string is not multiple of blocksize")
	}
	
	lastPad := text[len(text)-1]
	padLen := int(lastPad)
	trimmed := text
	
    for trimmed[len(trimmed)-1] == lastPad {
        trimmed = bytes.TrimSuffix(trimmed, []byte{lastPad})
    }

	if len(trimmed) != len(text)-padLen {
		return nil, fmt.Errorf("invalid padding")
	}
	return trimmed, nil
}
