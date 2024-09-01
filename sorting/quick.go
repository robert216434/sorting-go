package sortinggo

import "golang.org/x/exp/constraints"

func QuickSort[V constraints.Ordered](slice []V) {
	sort(slice, 0, len(slice)-1)
}

func sort[V constraints.Ordered](arr []V, low, high int) []V {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = sort(arr, low, p-1)
		arr = sort(arr, p+1, high)
	}
	return arr
}

func partition[V constraints.Ordered](arr []V, low, high int) ([]V, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
