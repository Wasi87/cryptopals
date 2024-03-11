package set1

import "sort"

type Pair struct {
	Key      byte
	XordBytes []byte
}

type Result struct {
	Score     float64
	XordString string
}

type ResultSlice []Result

func (rs ResultSlice) Len() int {
	return len(rs)
}

func (rs ResultSlice) Less(i, j int) bool {
	return rs[i].Score > rs[j].Score
}

func (rs ResultSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// xor - str^所有單一字符 ? str^單一字符
// 返回 map 對應ascii / xor結果
func XorEverySingleChar(str string) []Pair {
	bytes := HexToBytes(str)
	var decryptedPairs []Pair

	for i := 0; i < 256; i++ {
		var xordBytes []byte
		for _, b := range bytes {
			xordByte := b ^ byte(i)
			xordBytes = append(xordBytes, xordByte)
		}
		decryptedPairs = append(decryptedPairs, Pair{Key: byte(i), XordBytes: xordBytes})
	}
	return decryptedPairs
}

// https://en.wikipedia.org/wiki/Letter_frequency (sum 100.12%) 
func score(p []Pair) []Result {
	textsFrequency := map[byte]float64{
		'a': 8.2,
		'b': 1.5,
		'c': 2.8,
		'd': 4.3,
		'e': 12.7,
		'f': 2.2,
		'g': 2.0,
		'h': 6.1,
		'i': 7.0,
		'j': 0.15,
		'k': 0.77,
		'l': 4.0,
		'm': 2.4,
		'n': 6.7,
		'o': 7.5,
		'p': 1.9,
		'q': 0.095,
		'r': 6.0,
		's': 6.3,
		't': 9.1,
		'u': 2.8,
		'v': 0.98,
		'w': 2.4,
		'x': 0.15,
		'y': 2.0,
		'z': 0.074,
		' ': 15,
	}
	var results []Result
	var result Result
	for _, pair := range p {
		var totalScore float64
		for _, xordByte := range pair.XordBytes {
			frequency, found := textsFrequency[xordByte]
			if found {
				totalScore += frequency
			}
			if xordByte+32 >= 97 && xordByte+32 <= 122 {
				totalScore += frequency
			}
		}
		xordString := string(pair.XordBytes)

		result = Result{
			Score:      totalScore,
			XordString: xordString,
		}
		results = append(results, result)
	}

	return results
}

func GetTopFiveScore(input string) []Result {
	decryptedPairs := XorEverySingleChar(input)
	results := score(decryptedPairs)
	sort.Sort(ResultSlice(results))

	return results[:5]
}
