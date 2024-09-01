package sortinggo

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// Uses pattern-defeating quicksort (pdqsort)
func SortFuncSlices[V constraints.Ordered](slice []V) {
	slices.SortFunc(slice, func(i, j V) int {
		switch {
		case i < j:
			return -1
		case i > j:
			return 1
		default:
			return 0
		}
	})
}
