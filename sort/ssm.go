package sort

type ints []int

func (x ints) swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func Ssm(a ints) []int {
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
		// current := a[i]
		// a[i] = min
		// a[minIndex] = current
		a.swap(i, minIndex)
	}

	return a
}
