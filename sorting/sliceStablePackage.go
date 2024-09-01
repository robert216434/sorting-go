package sortinggo

import (
	sortPackage "sort"

	"golang.org/x/exp/constraints"
)

// Uses insertion sort
func SliceStablePackage[V constraints.Ordered](slice []V) {
	sortPackage.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
