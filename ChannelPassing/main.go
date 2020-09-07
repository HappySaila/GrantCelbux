package main

import (
	"fmt"
	"time"
)

func main() {
	//initialize
	t := make(chan int)
	go func() {
		t <- 0
	}()
	go ChannelFunnel(t)
	time.Sleep(time.Minute)
}

func ChannelFunnel(x chan int) {
	val := <- x
	fmt.Println(val)
	if val > 4 {
		fmt.Println("Done")
		return
	}
	y := make(chan int)
	go ChannelFunnel(y)
	y <- val+1
}


