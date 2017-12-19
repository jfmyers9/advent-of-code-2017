package helpers

import "strconv"

func MustConvertToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}
