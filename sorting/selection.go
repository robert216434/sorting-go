package sortinggo

import "golang.org/x/exp/constraints"

func SelectionSort[V constraints.Ordered](slice []V) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[min_idx] {
				min_idx = j
			}
		}

		temp := slice[min_idx]
		slice[min_idx] = slice[i]
		slice[i] = temp
	}
}
