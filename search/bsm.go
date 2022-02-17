package search

import (
	"fmt"
	"math"
	"time"

	"github.com/frappuccino2316/go-algorithm/cerror"
)

func Bsm(a []int, b int) (int, error) {
	time.Sleep(time.Second)
	fmt.Printf("target: %v\n", a)
	fmt.Printf("search: %v\n", b)

	length := len(a)
	point := length / 2
	newPoint := point
	upperLimit := length
	lowerLimit := 0

	for i := 0; i < int(math.Log2(float64(length)))+1; i++ {
		if length == 1 {
			if a[newPoint] == b {
				return newPoint, nil
			} else {
				return -1, cerror.RaiseError()
			}
		}

		if a[newPoint] == b {
			return newPoint, nil
		} else if a[newPoint] < b {
			lowerLimit = newPoint
			newPoint = lowerLimit + (upperLimit-lowerLimit)/2
		} else {
			upperLimit = newPoint
			newPoint = lowerLimit + (upperLimit-lowerLimit)/2
		}
	}
	return -1, cerror.RaiseError()
}
