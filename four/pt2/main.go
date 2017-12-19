package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
		sortedElem := sortString(elem)
		if _, ok := words[sortedElem]; ok {
			return false
		}
		words[sortedElem] = struct{}{}
	}
	return true
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
