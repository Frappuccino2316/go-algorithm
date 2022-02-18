package search

// import (
// 	"fmt"
// 	"time"

// 	"github.com/frappuccino2316/go-algorithm/cerror"
// )

func Hashing(a []int, b int) []int {
	length := len(a)
	var h []int
	for _, n := range a {
		k := n % (length + 1)
		if h[k] == 0 {
			h[k] = n
		} else {
			for i := k + 1; i == (k + length); i++ {
				if h[i] == 0 {
					h[i] = n
				}
				continue
			}
		}
	}
	return h
}
