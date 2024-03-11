package set1

import (
	"bufio"
	"cmp"
	"os"
	"slices"
)

func DetectSingleCharXor(f string) []Result{
	file, _ := os.Open(f)
	defer file.Close()

	var results []Result
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := GetTopFiveScore(line)

		results = append(results, result...)
	}
	
	slices.SortFunc(results, func(a, b Result) int {
		n := cmp.Compare(b.Score, a.Score)  
		return n
	})
	return results
}