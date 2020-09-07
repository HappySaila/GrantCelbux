package main

import (
	"fmt"
	"math/rand"
	"time"
)

var p = fmt.Println
var n = 10000
var thresh = n/4

func main() {
	//Create an array to compute
	//Generate seed to get a different array everytime
	rand.Seed(int64(time.Now().Nanosecond()))
	var y []int64
	p("Sandbox > Playground")
	for i := 0; i < n; i++ {
		y = append(y, rand.Int63n(10) + int64(1))
	}
	//sequential
	start := time.Now()
	p("Total: ", SequencialSum(y))
	p("Sequential Duration: ", time.Since(start))

	//parallel
	start = time.Now()
	total := make(chan int64)
	go ParrellelSum(y, thresh, 0, len(y), total)
	p("Total: ", <-total)
	p("Parallel Duration: ", time.Since(start))
}

func SequencialSum(arr []int64) int64 {
	total := int64(0)
	for _, v := range arr {
		total = total + int64(v)
		time.Sleep(time.Millisecond)
	}
	return total
}

func ParrellelSum(arr []int64, thresh int, start int, end int, result chan int64) {
	if end - start <= thresh {
		result <- SequencialSum(arr[start:end])
		return
	}

	subResult := make(chan int64)
	subResult2 := make(chan int64)

	go func(arr []int64, thresh int, start int, end int) {
		p("Goroutine spawned")
		ParrellelSum(arr, thresh, start, (start+end)/2, subResult)
	}(arr, thresh, start, end)

	go func(arr []int64, thresh int, start int, end int) {
		p("Goroutine spawned")
		ParrellelSum(arr, thresh, (start+end)/2, end, subResult2)
	}(arr, thresh, start, end)
	val := <-subResult + <-subResult2
	result <- val
}


