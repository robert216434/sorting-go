package sortinggo_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"

	sortinggo "github.com/robert216434/sorting-go/sorting"
)

// To run check https://github.com/robert216434/sorting-go?tab=readme-ov-file#running-benchmark

const sliceSize = 10_000
const maxValue = 1_000_000_000

var sliceToSort = []int{}

func getSliceCopy(slice []int) []int {
	return append([]int(nil), slice...)
}

func init() {
	start := time.Now()
	for i := 0; i < sliceSize; i++ {
		genInt, _ := rand.Int(rand.Reader, big.NewInt(maxValue+1))
		sliceToSort = append(sliceToSort, int(genInt.Int64()))
	}
	end_time := time.Since(start)
	fmt.Printf("generated slice with size %d, max value %d in %d nanoseconds / %d milliseconds\n", sliceSize, maxValue, end_time.Nanoseconds(), end_time.Milliseconds())
}

func Benchmark_MergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.MergeSort(getSliceCopy(sliceToSort))
	}
}

func Benchmark_NaiveSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.NaiveSort(getSliceCopy(sliceToSort))
	}
}

func Benchmark_QuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.QuickSort(getSliceCopy(sliceToSort))
	}
}

func Benchmark_SlicePackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.SlicePackage(getSliceCopy(sliceToSort))
	}
}

func Benchmark_SliceStablePackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.SliceStablePackage(getSliceCopy(sliceToSort))
	}
}

func Benchmark_SortFuncSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortinggo.SortFuncSlices(getSliceCopy(sliceToSort))
	}
}
