package sortinggo_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"

	sortinggo "github.com/robert216434/sorting-go/sorting"
)

// To run:
// go test -bench . -benchmem

const sliceSize = 10_000
const maxValue = 1_000_000_000

var sliceToSort = []int{}

func init() {
	start := time.Now()
	for i := 0; i < sliceSize; i++ {
		genInt, _ := rand.Int(rand.Reader, big.NewInt(maxValue+1))
		sliceToSort = append(sliceToSort, int(genInt.Int64()))
	}
	end_time := time.Since(start)
	fmt.Printf("generated slice with size %d, max value %d in %d nanoseconds\n", sliceSize, maxValue, end_time.Nanoseconds())
}

func Benchmark_MergeSort(b *testing.B) {
	sortinggo.MergeSort(append([]int(nil), sliceToSort...))
}

func Benchmark_NaiveSort(b *testing.B) {
	sortinggo.NaiveSort(append([]int(nil), sliceToSort...))
}
