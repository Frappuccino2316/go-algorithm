package sort

func Bubble(a []int) []int {
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
		current := a[i]
		a[i] = min
		a[index] = current
	}

	return a
}
