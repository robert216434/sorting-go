package sortinggo

import "golang.org/x/exp/constraints"

func BubbleSort[V constraints.Ordered](slice []V) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		exit := true
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp

				exit = false
			}
		}

		if exit {
			break
		}
	}
}
