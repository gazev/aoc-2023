package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	fmt.Printf("result 1st: %d\n", first())
	fmt.Printf("result 2nd: %d\n", second())
}

func first() int {
	fd, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("error opening input file -> %s\n", err)
	}

	res := 0
	r := bufio.NewReader(io.Reader(fd))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		first, last := 0, 0
		for _, r := range line {
			if !unicode.IsNumber(r) {
				continue
			}

			if first == 0 {
				first = int(r - '0')
			}
			last = int(r - '0')
		}
		res += first*10 + last
	}
	return res
}

var numbersStr = map[rune][]string{
	'o': {"one"},
	't': {"two", "three"},
	'f': {"four", "five"},
	's': {"six", "seven"},
	'e': {"eight"},
	'n': {"nine"},
}

var numbersVal = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func second() int {
	fd, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("error opening input file -> %s\n", err)
	}

	res := 0
	r := bufio.NewReader(io.Reader(fd))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		lineLen := len(line)
		first, last := 0, 0
		for i, r := range line {
			if unicode.IsNumber(r) {
				if first == 0 {
					first = int(r - '0')
				}
				last = int(r - '0')
				continue
			}

			numbersSlice, inMap := numbersStr[r]
			if !inMap {
				continue
			}

			for _, numberString := range numbersSlice {
				offset := len(numberString)

				if i+offset > lineLen {
					continue
				}

				if line[i:i+offset] != numberString {
					continue
				}

				if first == 0 {
					first = numbersVal[line[i:i+offset]]
				}
				last = numbersVal[line[i:i+offset]]
			}
		}
		res += first*10 + last
	}
	return res
}
