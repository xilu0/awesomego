package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-1 square root is", result)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return f*f, nil
	
}
