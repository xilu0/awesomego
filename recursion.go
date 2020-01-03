package main

import "fmt"

func main() {
	var i int = 3
	fmt.Printf("%d 的阶乘是 %d", i, Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if (n > 0){
		result = n * Factorial(n-1)
		fmt.Println("11111111")
		return result
	}
	return 1
}