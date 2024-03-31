package set1

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"log"
	"os"

	_ "github.com/gofiber/fiber/v2/log"
)

// bytes to chunk
func BytesToChunks(line []byte, blockSize int) [][]byte {
	chunks := make([][]byte, len(line)/blockSize)
	for i := 0; i < len(line); i += blockSize {
		end := i + blockSize
		if end > len(line) {
			end = len(line)
		}
		chunks[i/blockSize] = line[i:end]
	}
	return chunks
}

func DetectECB(filePath string)(int, int){
	file, _ := os.Open(filePath)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	blockSize := 16
	maxCount := 0
	maxRow := 0
	currentRow := 0
	
	for scanner.Scan(){
		currentRow++
		decodedLine, err := hex.DecodeString(scanner.Text())
		if err != nil{
			log.Fatal(err)
		}
		chunks := BytesToChunks(decodedLine, blockSize)

		for i := 0 ; i < len(chunks)-1 ; i++ {
			count := 0
			for j := i+1 ; j < len(chunks); j++ {
				if bytes.Equal(chunks[i], chunks[j]) {
					count++
				}
			}
			if count > maxCount {
				maxCount = count
				maxRow = currentRow
			}
		}
	}
	return maxRow, maxCount
}
