package set1

func RepeatingKeyXor(key string, input string) []byte {
	keyBytes := []byte(key)
	inputBytes := []byte(input)
	var xorResults []byte
	for i := 0; i < len(inputBytes); i++ {
		xorResult := inputBytes[i] ^ keyBytes[i%len(keyBytes)]
		xorResults = append(xorResults, xorResult)
	}
	
	return xorResults
}