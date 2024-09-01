package sortinggo

import "slices"

func RadixSort(slice []int) {
	max := slices.Max(slice)
	sliceSize := len(slice)
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(slice, sliceSize, exp)
	}
}

func countSort(slice []int, sliceSize, exp int) {
	output := make([]int, sliceSize)
	count := make([]int, 10)

	for i := 0; i < sliceSize; i++ {
		index := (slice[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := sliceSize - 1; i >= 0; i-- {
		cIdx := (slice[i] / exp) % 10
		oIdx := count[cIdx] - 1
		output[oIdx] = slice[i]
		count[cIdx]--
	}

	for i := 0; i < sliceSize; i++ {
		slice[i] = output[i]
	}
}
