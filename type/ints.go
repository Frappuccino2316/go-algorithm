package ints

type Ints []int

func (x Ints) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
