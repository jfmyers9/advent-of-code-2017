package main

import (
	"flag"
	"fmt"
)

var input = flag.Int("input", 0, "Input to the puzzle")

type point struct {
	x int
	y int
}

func main() {
	flag.Parse()

	values := map[point]int{}
	values[point{x: 0, y: 0}] = 1
	max := 1
	for {
		for i := 1; i <= max*2; i++ {
			n := calculatePoint(max, -max+i, values)
			if n > *input {
				fmt.Printf("Next: %d\n", n)
				return
			}
		}

		for i := 1; i <= max*2; i++ {
			n := calculatePoint(max-i, max, values)
			if n > *input {
				fmt.Printf("Next: %d\n", n)
				return
			}
		}

		for i := 1; i <= max*2; i++ {
			n := calculatePoint(-max, max-i, values)
			if n > *input {
				fmt.Printf("Next: %d\n", n)
				return
			}
		}

		for i := 1; i <= max*2; i++ {
			n := calculatePoint(-max+i, -max, values)
			if n > *input {
				fmt.Printf("Next: %d\n", n)
				return
			}
		}

		max++
	}
}

func calculatePoint(i, j int, values map[point]int) int {
	sum := 0
	if x, ok := values[point{i - 1, j}]; ok {
		sum += x
	}
	if x, ok := values[point{i - 1, j - 1}]; ok {
		sum += x
	}
	if x, ok := values[point{i, j - 1}]; ok {
		sum += x
	}
	if x, ok := values[point{i + 1, j - 1}]; ok {
		sum += x
	}
	if x, ok := values[point{i + 1, j}]; ok {
		sum += x
	}
	if x, ok := values[point{i + 1, j + 1}]; ok {
		sum += x
	}
	if x, ok := values[point{i, j + 1}]; ok {
		sum += x
	}
	if x, ok := values[point{i - 1, j + 1}]; ok {
		sum += x
	}

	values[point{i, j}] = sum
	return sum
}
