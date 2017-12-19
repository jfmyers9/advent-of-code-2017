package main

import (
	"flag"
	"fmt"
)

var input = flag.Int("input", 0, "Input to the puzzle")

func main() {
	flag.Parse()

	m := 1
	i := 0
	for {
		n := m + (8 * i)
		if n > *input {
			break
		}

		m = n
		i++
	}

	location := (*input - m) % (i * 8)

	offset := (i * 8) / 4
	relativeLocation := location % offset

	halfwayPoint := (i * 8) / 4 / 2

	moves := 0
	for {
		if relativeLocation > (i*8)/4/2 {
			moves++
			relativeLocation--
		}

		if relativeLocation < (i*8)/4/2 {
			moves++
			relativeLocation++
		}

		if relativeLocation == halfwayPoint {
			moves += i
			fmt.Printf("Moves: %d\n", moves)
			break
		}
	}
}
