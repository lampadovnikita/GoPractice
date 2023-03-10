package main

import (
	"fmt"
	"runtime"
	"time"
)

const N = 40000000

func main() {
	myMap := make(map[int]*int)

	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}

	start := time.Now()
	runtime.GC()
	duration := time.Since(start)

	fmt.Println("It took GC", duration, "to finish")

	_ = myMap[0]
}
