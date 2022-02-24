package sort

import (
	"time"
)

func Sim(a []int) []int {
	var new []int

	for i, n := range a {
		// fmt.Println(new)
		if i == 0 {
			new = append(new, a[i])
		} else {
			for j, m := range new {
				time.Sleep(time.Second)

				firstHalf := append([]int{}, new[:j]...)
				latterHalf := append([]int{}, new[j:]...)

				if m > n {
					tmp := append(firstHalf, n)
					new = append(tmp, latterHalf...)
					break
				} else if m == n {
					tmp := append(firstHalf, n)
					new = append(tmp, latterHalf...)
				} else if m < n {
					if j == len(new)-1 {
						new = append(new, n)
					}
					continue
				}
			}
		}
	}
	return new
}
