package main

import (
	"fmt"
	"time"
)

func main() {
	go say("word")
	say("hello")
}

func say(s string) {
	for i:=0; i<10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
