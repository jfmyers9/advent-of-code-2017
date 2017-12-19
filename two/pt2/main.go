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

		sum += findEvenDivisor(elements)
	}

	fmt.Printf("Checksum: %d\n", sum)
}

func findEvenDivisor(elements []string) int {
	for i := 0; i < len(elements); i++ {
		x := mustConvertToInt(elements[i])
		for j := i + 1; j < len(elements); j++ {
			y := mustConvertToInt(elements[j])

			if x%y == 0 {
				fmt.Printf("Found Even Divisor: %d / %d\n", x, y)
				return x / y
			}

			if y%x == 0 {
				fmt.Printf("Found Even Divisor: %d / %d\n", y, x)
				return y / x
			}
		}
	}

	panic("There is no even Divisor on this line")
}

func mustConvertToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
