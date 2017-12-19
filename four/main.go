package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
		if validPassphrase(line) {
			sum++
		}
	}

	fmt.Printf("Num Valid: %d\n", sum)
}

func validPassphrase(s string) bool {
	elements := strings.Split(s, " ")
	words := map[string]struct{}{}
	for _, elem := range elements {
		if _, ok := words[elem]; ok {
			return false
		}

		words[elem] = struct{}{}
	}
	return true
}
