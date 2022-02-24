package main

import (
	"fmt"
	"time"

	// "github.com/frappuccino2316/go-algorithm/cerror"
	// "github.com/frappuccino2316/go-algorithm/search"
	"github.com/frappuccino2316/go-algorithm/sort"
)

// 単純挿入法 Simple insertion method
// func main() {
// 	a := []int{3, 45, 13, 84, 47, 65, 123, 542, 85, 99, 12, 1, 45, 26, 75, 23}
// 	now := time.Now()
// 	b := sort.Sim(a)
// 	fmt.Printf("秒数: %vms\n", time.Since(now).Milliseconds())

// 	fmt.Println(b)
// }

// 単純交換法　Bubble sort
func main() {
	a := []int{3, 45, 13, 84, 47, 65, 123, 542, 85, 99, 12, 1, 45, 26, 75, 23}
	now := time.Now()
	b := sort.Bubble(a)
	fmt.Printf("秒数: %vms\n", time.Since(now).Milliseconds())

	fmt.Println("result: ", b)
}

// 単純選択法 Simple selection method
// func main() {
// 	a := []int{3, 45, 13, 84, 47, 65, 123, 542, 85, 99, 12, 1, 45, 26, 75, 23}
// 	now := time.Now()
// 	b := sort.Ssm(a)
// 	fmt.Printf("秒数: %vms\n", time.Since(now).Milliseconds())

// 	fmt.Printf("result: %v\n", b)
// }

// ハッシュ探索法 hashing method
// func main() {
// 	a := []int{1, 8, 11, 15, 17, 18, 21, 37}
// 	h := search.Hashing(a)

// 	i, err := search.Hm(h, 15)

// 	if err != nil {
// 		e, ok := err.(*cerror.NothingError)
// 		if ok {
// 			fmt.Println(e.Message)
// 		}
// 	} else {
// 		fmt.Printf("index: %v\n", i)
// 		fmt.Printf("value: %v\n", h[i])
// 	}
// }

// 二分探索法 Binary search method
// func main() {
// 	a := []int{1, 2, 4, 6, 11, 15, 17, 18, 21, 50, 55, 67, 73, 99, 100}

// 	i, err := search.Bsm(a, 101)
// 	if err != nil {
// 		e, ok := err.(*cerror.NothingError)
// 		if ok {
// 			fmt.Println(e.Message)
// 		}
// 	} else {
// 		fmt.Printf("index: %v\n", i)
// 		fmt.Printf("value: %v\n", a[i])
// 	}
// }

// 線形探索法 Linear search method
// func main() {
// 	a := []int{3, 2, 6, 9, 6, 2, 5, 6, 12, 44, 89, 125, 45, 1, 435, 2, 534, 23, 5, 2, 64}

// 	i, err := search.Lsm(a, 1)
// 	if err != nil {
// 		e, ok := err.(*cerror.NothingError)
// 		if ok {
// 			fmt.Println(e.Message)
// 		}
// 	} else {
// 		fmt.Printf("index: %v\n", i)
// 		fmt.Printf("value: %v\n", a[i])
// 	}
// }
