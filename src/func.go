package main

import (
	"fmt"
)

func main() {
	bestFinish := bestFinishes(10, 20, 30, 1, 2)

	fmt.Println(bestFinish)
}

func bestFinishes(modules ...int) int {
	best := modules[0]
	for _, i := range modules {
		if i < best {
			best = i
		}
	}

	return best
}