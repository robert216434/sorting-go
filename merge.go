package sortinggo

import "golang.org/x/exp/constraints"

func MergeSort[V constraints.Ordered](slice []V) []V {
	if len(slice) <= 1 {
		return slice
	}
	middle := len(slice) / 2
	left := MergeSort(slice[:middle])
	right := MergeSort(slice[middle:])
	return merge(left, right)
}

func merge[V constraints.Ordered](left, right []V) []V {
	result := make([]V, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}
