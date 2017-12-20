package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	for s.Scan() {
		line := s.Text()
		elements := strings.Split(line, "\t")
		numbers := convertElementsToInts(elements)
		fmt.Printf("Loop Size: %d\n", performRedistribution(numbers))
		return
	}
}

func convertElementsToInts(values []string) []int {
	result := make([]int, len(values))
	for i, s := range values {
		result[i] = helpers.MustConvertToInt(s)
	}

	return result
}

func convertElementsToString(values []int) string {
	result := make([]string, len(values))
	for i, n := range values {
		result[i] = strconv.Itoa(n)
	}

	return strings.Join(result, ",")
}

func performRedistribution(numbers []int) int {
	seen := map[string]int{}
	seen[convertElementsToString(numbers)] = 0
	moves := 0
	for {
		moves++

		maxPos := 0
		max := 0
		for i, n := range numbers {
			if n > max {
				max = n
				maxPos = i
			}
		}

		numbers[maxPos] = 0
		position := maxPos + 1
		for i := max; i > 0; i-- {
			if position >= len(numbers) {
				position = 0
			}

			numbers[position]++
			position++
		}

		key := convertElementsToString(numbers)
		if prevMove, ok := seen[key]; ok {
			return moves - prevMove
		}

		seen[key] = moves
	}
}
