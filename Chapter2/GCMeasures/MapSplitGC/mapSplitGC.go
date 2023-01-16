package main

import "runtime"

const N = 40000000

func main() {
	split := make([]map[int]int, 200)

	for i := range split {
		split[i] = make(map[int]int)
	}

	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}

	runtime.GC()
	_ = split[0][0]
}
