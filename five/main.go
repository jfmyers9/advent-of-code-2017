package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jfmyers9/advent-of-code-2017/helpers"
)

var input = flag.String("input", "", "Input to the puzzle")

func main() {
	flag.Parse()

	if *input == "" {
		panic("provide an input file")
	}

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var maze []int
	for s.Scan() {
		line := s.Text()
		maze = append(maze, helpers.MustConvertToInt(line))
	}

	position := 0
	moves := 0
	for {
		if position < 0 || position >= len(maze) {
			fmt.Printf("Moves: %d\n", moves)
			return
		}

		n := maze[position]
		maze[position] = n + 1
		position += n
		moves++
	}
}
