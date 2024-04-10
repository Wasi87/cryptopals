package set2

import (
	"bytes"
	set1 "cryptopals/set-1"
	"cryptopals/util"
	"math/rand"
	"slices"
	"sync"
	"time"
)

var (
	key14 []byte
	randPref []byte
	once14 sync.Once
)

func init(){
	once14.Do(func() {
		key14 = GenAESKey(16)
		randPref = GenRandPref()
	})
}

func GenRandPref() []byte {
	randLen := rand.Intn(128)
	pref := make([]byte, randLen)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	r.Read(pref)
	return pref
}

func RandPrefixOracle(plainText, key []byte) []byte {
	text := set1.ReadFile("texts/ch12.txt")
	unknown := set1.DecodeBase64(text)

	plainText = append(randPref, plainText...)
	plainText = append(plainText, unknown...)
	encrypted := util.AesEncryptECB(key14, plainText)
	return encrypted
}

func Chal14() []byte {

	preLength := 0

	var bs int

	for i := 0;; i++ {
		// find blockSize
		input := bytes.Repeat([]byte("A"), i)
		ct := RandPrefixOracle(input, key14)
		if len(ct) > preLength && preLength != 0 {
			bs = len(ct) - preLength
			break
		} 
		preLength = len(ct) 
	}

	// verify ECB
	count := util.DetectECB(RandPrefixOracle(bytes.Repeat([]byte("A"), 4*bs), key14), bs)
	if count == 0 {
	panic("Not ECB mode")
	}

	// find length of random number prefix
	inpIdx := -1
	var lenPrefix int
	var lenPt int
	var filln int
	foundIdx := false
	for i := 0; ; i++ {
		input := bytes.Repeat([]byte("A"), i)
		inp1, inp2 := []byte("A"), []byte("B")
		cp1, cp2 := append(input, inp1...), append(input, inp2...)
		ct1, ct2 := RandPrefixOracle(cp1, key14), RandPrefixOracle(cp2, key14)
		chunks1, chunks2 := set1.BytesToChunks(ct1, bs), set1.BytesToChunks(ct2, bs)
		if !foundIdx && !slices.Equal(chunks1[i], chunks2[i]) {
			inpIdx = i
			foundIdx = true
			}	
		if inpIdx != -1 && bytes.Equal(chunks1[inpIdx], chunks2[inpIdx]) {
			filln = i
			lenPrefix = (inpIdx+1)*bs - i
			lenPt = len(ct1) - lenPrefix - i 
			break
		}
	}

	// crack the plaintext
	toFill := bytes.Repeat([]byte("A"), filln)
	var ans []byte
	for k := (inpIdx+1)*bs  ; k < (inpIdx+1)*bs+lenPt; k+=bs {
		for i := bs-1; i >= 0; i-- {			
			input := bytes.Repeat([]byte("A"), i)
			input = append(toFill, input...)
			ct := RandPrefixOracle(input , key14)
			for j := 0; j < 255; j++ {
				guess := bytes.Repeat([]byte("A"),filln)
				if i > 0 {
					re := bytes.Repeat([]byte("A"), i)
					guess = append(guess, re...)
				} 
				if len(ans) != 0 {
					guess = append(guess, ans...)
				}
				guess = append(guess, byte(j))		
				ctGuess := RandPrefixOracle(guess, key14)
				if slices.Equal(ct[k:k+bs], ctGuess[k:k+bs]) {
					ans = append(ans, byte(j))
					break
				}
			}
		}
	}
	return ans
}