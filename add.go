package main

import "fmt"

func main() {
     fmt.Println("1 + 到10的结果是", andadd(10))
}

func andadd(i int) (result int) {
	if(i>0){
		result = i+andadd(i-1)
		return
	} else {
		return 0
	}
}
