package main

import (
	"fmt"
	"go-algorithm/cerror"
	"go-algorithm/search"
)

func main() {
	a := []int{3, 2, 6, 9, 1}

	i, err := search.Lsm(a, 1)
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
