package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i int
	j int
}

const N = 40000000

func main() {
	var structure []data

	fmt.Printf("n = %v, type = %T", N, N)

	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	_ = structure
}
