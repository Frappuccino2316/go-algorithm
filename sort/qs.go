package sort

import (
	t "github.com/frappuccino2316/go-algorithm/type"
)

func Quicksort(a t.Ints) []int {
	if len(a) < 2 {
		return a
	}

	pivot := a[0]

	left := []int{}
	right := []int{}

	for _, v := range a[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	left = Quicksort(left)
	right = Quicksort(right)

	result := append(left, pivot)
	result = append(result, right...)

	return result
}
