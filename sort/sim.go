package sort

// ググった実装
// 配列の要素をずらして開けたindexに突っ込む
func Sim(a []int) []int {
	var i, j, target int

	for j = 1; j < len(a); j++ {
		target = a[j]
		i = j - 1

		for i >= 0 && a[i] > target {
			a[i+1] = a[i]
			i--
		}

		a[i+1] = target
	}

	return a
}

// 自力実装
// func Sim(a []int) []int {
// 	var new []int

// 	for i, n := range a {
// 		if i == 0 {
// 			new = append(new, a[i])
// 		} else {
// 			for j, m := range new {
// 				time.Sleep(time.Second)

// 				firstHalf := append([]int{}, new[:j]...)
// 				latterHalf := append([]int{}, new[j:]...)

// 				if m > n {
// 					tmp := append(firstHalf, n)
// 					new = append(tmp, latterHalf...)
// 					break
// 				} else if m == n {
// 					tmp := append(firstHalf, n)
// 					new = append(tmp, latterHalf...)
// 				} else if m < n {
// 					if j == len(new)-1 {
// 						new = append(new, n)
// 					}
// 					continue
// 				}
// 			}
// 		}
// 	}
// 	return new
// }
