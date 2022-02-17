package main

import (
	"fmt"

	"github.com/frappuccino2316/go-algorithm/cerror"
	"github.com/frappuccino2316/go-algorithm/search"
)

// 二分探索法 Binary search method
func main() {
	a := []int{1, 2, 4, 6, 11, 15, 17, 18, 21, 50, 55, 67, 73, 99, 100}

	i, err := search.Bsm(a, 101)
	if err != nil {
		e, ok := err.(*cerror.NothingError)
		if ok {
			fmt.Println(e.Message)
		}
	} else {
		fmt.Printf("index: %v\n", i)
		fmt.Printf("value: %v\n", a[i])
	}
}

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
