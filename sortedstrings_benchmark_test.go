package handy

import (
	"bufio"
	"math/rand"
	"os"
	"sort"
	"testing"
)

const (
	dictFile = "testdata/words"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func BenchmarkBinarySearch(b *testing.B) {
	words, err := readLines(dictFile)
	if err != nil {
		b.Fatal(err)
	}
	dict := SortedStrings(words)
	size := len(words)
	for i := 0; i < b.N; i++ {
		r := rand.Int31n(int32(size))
		word := dict[r]
		j := sort.SearchStrings(dict, word)
		if int32(j) != r {
			b.Fatalf("fail random search at iteration %d, word %s", i, word)
		}

	}
}

func BenchmarkBuiltinBinarySearch(b *testing.B) {
	words, err := readLines(dictFile)
	if err != nil {
		b.Fatal(err)
	}
	dict := SortedStrings(words)
	size := len(words)
	for i := 0; i < b.N; i++ {
		r := rand.Int31n(int32(size))
		word := dict[r]
		j := dict.BinarySearch(word)
		if int32(j) != r {
			b.Fatalf("fail random search at iteration %d, word %s", i, word)
		}

	}
}
