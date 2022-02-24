package sort

import (
	. "github.com/frappuccino2316/go-algorithm/type"
)

func Ssm(a Ints) []int {
	length := len(a)

	for i := 0; i < length; i++ {
		min := a[i]
		minIndex := i

		for j := i; j < length; j++ {
			if j == i {
				continue
			}

			if a[j] < min {
				min = a[j]
				minIndex = j
			}
		}
		a.Swap(i, minIndex)
	}

	return a
}
