package set1

import (
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func HammingDistance(str1, str2 []byte) (int, error) {
	if len(str1) != len(str2) {
		return 0, fmt.Errorf("two strings must have the same length")
	}
	distance := 0
	for i := 0; i < len(str1); i++ {
		binStr1 := fmt.Sprintf("%08b", str1[i])
		binStr2 := fmt.Sprintf("%08b", str2[i])

		for j := 0; j < len(binStr1); j++ {
			if binStr1[j] != binStr2[j] {
				distance++
			}
		}
	}
	return distance, nil
}


// [1,20,3,43] / 兩兩距離 3, 37, 2 / 平均 (3+37+2)/3
func FindKeyLength(enc []byte) (int, float64) {
	// chunkSize := 4 doesn't work!!!!!!!
	iter := 0
	var avgDist float64
	var key int
	minAvgDist := math.MaxFloat64
	
	for keySize := 2; keySize <= 40; keySize++ {
		if len(enc) < 4*keySize {
			log.Fatal("length of encrypted is less than 4 x keysize")
			break
		}

		sumDist := 0
		for j := 0; j+2*keySize < len(enc); j+=keySize {
			iter ++
			dist, err := HammingDistance(enc[j:j+keySize], enc[j+keySize:j+2*keySize])
			if err != nil {
				fmt.Println("Error calculating Hamming distance:", err)
			}
			sumDist += dist
			avgDist = float64(sumDist) / float64(iter) 
		}

		if avgDist < minAvgDist {
			minAvgDist = avgDist
			key = keySize
		}
	}
	return key, minAvgDist
}


func Base64ToBinary(str string)([]byte, error){
	binData, err := base64.StdEncoding.DecodeString(str)
	if err != nil{
		return nil, err
	}
	return binData, nil
}

func FindKeyChar(enc []byte, keySize int) []byte {
	blocks := make([][]byte, (len(enc)/keySize)+1)

	for i := 0; i < len(enc); i += keySize {
		end := i + keySize
		if end > len(enc) {
			end = len(enc)
		}
		blocks[i/keySize] = append(blocks[i/keySize], enc[i:end]...)
	}

	// blocks [ []byte29, 29, 29...]
	transBlocks := make([][]byte, keySize)
    for i := range transBlocks {
        transBlocks[i] = make([]byte, len(blocks))
        for j := range transBlocks[i] {
            if len(blocks[j]) > i {
                transBlocks[i][j] = blocks[j][i]
            }
        }
    }
	
	var key []byte
	for _, v := range transBlocks {
		decryptedPairs := XorEverySingleChar(v)
		results := Score(decryptedPairs)
		sort.Sort(ResultSlice(results))
		key = append(key, results[0].Key)
	}
	return key
}


func BreakRepeatingXor(filePath string) ([]byte, []byte) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	encrypted, err := Base64ToBinary(string(data))
	if err != nil {
		log.Fatal(err)
	}

	keyLen, _ := FindKeyLength(encrypted)
	key :=FindKeyChar(encrypted, keyLen)
	decrypted := RepeatingKeyXor(string(key), string(encrypted))

	return key, decrypted
}

