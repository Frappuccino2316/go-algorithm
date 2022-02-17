package search

import (
	"fmt"

	"github.com/frappuccino2316/go-algorithm/cerror"
)

func Lsm(a []int, b int) (int, error) {
	fmt.Printf("target: %v\n", a)
	fmt.Printf("search: %v\n", b)

	for i, n := range a {
		if n == b {
			fmt.Println("みつけた!")
			return i, nil
		}
		continue
	}
	return -1, cerror.RaiseError()
}
