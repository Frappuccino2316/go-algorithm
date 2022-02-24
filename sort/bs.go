package sort

import (
	. "github.com/frappuccino2316/go-algorithm/type"
)

func Bubble(a Ints) []int {
	length := len(a)

	for i := 0; i < length; i++ {
		index := length - 1
		min := a[index]

		for j := length - 1; j >= i; j-- {
			if a[j] < min {
				index = j
				min = a[j]
			}
		}
		a.Swap(i, index)
	}

	return a
}
