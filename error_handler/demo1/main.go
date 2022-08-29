package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func divFunc(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("b is 0")
		return 0, err
	}
	return a / b, nil
}

func main() {
	a := 1
	b := 0
	ret, err := divFunc(a, b)
	if err != nil {
		fmt.Printf("ret is %d, err is %+v", ret, err)
	}
}
