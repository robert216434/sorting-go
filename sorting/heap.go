package sortinggo

import "golang.org/x/exp/constraints"

func HeapSort[V constraints.Ordered](slice []V) {
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}

	for i := n - 1; i > 0; i-- {
		temp := slice[0]
		slice[0] = slice[i]
		slice[i] = temp

		heapify(slice, i, 0)
	}
}

func heapify[V constraints.Ordered](slice []V, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && slice[left] > slice[largest] {
		largest = left
	}

	if right < n && slice[right] > slice[largest] {
		largest = right
	}

	if largest != i {
		temp := slice[i]
		slice[i] = slice[largest]
		slice[largest] = temp

		heapify(slice, n, largest)
	}
}
