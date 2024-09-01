package sortinggo_test

import (
	"testing"

	sortinggo "github.com/robert216434/sorting-go/sorting"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	s := []int{78, 27, 34, 76, 32}
	sortinggo.SelectionSort(s)
	assert.Equal(t, []int{27, 32, 34, 76, 78}, s)
}
