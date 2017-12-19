package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var sum int
	for s.Scan() {
		line := s.Text()
		elements := strings.Split(line, "\t")

		if len(elements) <= 1 {
			panic("Invalid Spreadsheet")
		}

		n, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(err)
		}

		min := n
		max := n

		for i := 1; i < len(elements); i++ {
			n, err := strconv.Atoi(elements[i])
			if err != nil {
				panic(err)
			}

			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}

		sum += max - min
	}

	fmt.Printf("Checksum: %d\n", sum)
}
