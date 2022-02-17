package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 6, 9, 1}

	i, err := lsm(a, 1)
	if err != nil {
		e, ok := err.(*nothingError)
		if ok {
			fmt.Println(e.message)
		}
	} else {
		fmt.Printf("index: %v\n", i)
		fmt.Printf("value: %v\n", a[i])
	}
}

type nothingError struct {
	message string
}

func (e *nothingError) Error() string {
	return e.message
}

func raiseError() error {
	return &nothingError{message: "存在しませんでした"}
}

func lsm(a []int, b int) (int, error) {
	fmt.Printf("target: %v\n", a)
	fmt.Printf("search: %v\n", b)

	for i, n := range a {
		if n == b {
			fmt.Println("みつけた!")
			return i, nil
		}
		continue
	}
	return -1, raiseError()
}
