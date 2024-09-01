package sortinggo

import (
	sortPackage "sort"

	"golang.org/x/exp/constraints"
)

// Uses pattern-defeating quicksort (pdqsort)
func SliceSortPackage[V constraints.Ordered](slice []V) {
	sortPackage.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
