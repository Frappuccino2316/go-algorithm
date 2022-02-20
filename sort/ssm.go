package sort

func Ssm(a []int) []int {
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
		current := a[i]
		a[i] = min
		a[minIndex] = current
	}

	return a
}
