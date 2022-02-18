package search

import (
	"fmt"
	"time"

	"github.com/frappuccino2316/go-algorithm/cerror"
)

func Hashing(a []int) []int {
	length := len(a)
	h := make([]int, length*2)
	fmt.Println(len(h))

	for _, n := range a {
		k := n % (length + 1)
		if h[k] == 0 {
			h[k] = n
		} else {
			for i := k + 1; i < length*2; i++ {
				if h[i] == 0 {
					h[i] = n
					break
				}
			}
		}
	}
	return h
}

func Hm(h []int, b int) (int, error) {
	time.Sleep(time.Second)
	fmt.Printf("target: %v\n", h)
	fmt.Printf("search: %v\n", b)

	length := len(h)
	index := b % (length/2 + 1)

	for i := index; i < length; i++ {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("h[i]: %v\n", h[i])
		if h[i] == b {
			return i, nil
		}
	}
	return -1, cerror.RaiseError()
}
