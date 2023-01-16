package main

import "runtime"

const N = 40000000

func main() {
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}

	runtime.GC()
	_ = myMap[0]
}
