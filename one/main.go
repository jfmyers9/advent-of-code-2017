package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input = flag.String("input", "", "Input to the puzzle")

func main() {
	flag.Parse()

	if *input == "" {
		panic("provide an input file")
	}

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	trimmedData := strings.Trim(string(data), "\n")

	var sum int
	for i := 0; i < len(trimmedData); i++ {
		j := i + 1
		if j == len(trimmedData) {
			j = 0
		}

		a := trimmedData[i]
		b := trimmedData[j]

		if a == b {
			x, err := strconv.Atoi(string(a))
			if err != nil {
				panic(err)
			}

			sum += x
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
