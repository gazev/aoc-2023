package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	RED_CUBES   = 12
	GREEN_CUBES = 13
	BLUE_CUBES  = 14
)

func main() {
	fmt.Printf("result 1st: %d\n", first())
	fmt.Printf("result 2nd: %d\n", second())
}

func first() int {
	fd, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("error opening input file -> %s\n", err)
		return 0
	}
	defer fd.Close()

	reader := bufio.NewReader(io.Reader(fd))

	gameCount := 1
	res := 0
	for {
		lineB, err := reader.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0
		}

		line := strings.Split(string(lineB[:]), ":")[1]

		red, blue, green, offset := 0, 0, 0, 0
		inExpr := false
		currCountExpr := 0
		for i := 0; i+offset < len(line); i++ {
			char := line[i+offset]
			if char == ' ' {
				continue
			} else if char == ',' {
				inExpr = false
				continue
			} else if char == ';' {
				if red > RED_CUBES || blue > BLUE_CUBES || green > GREEN_CUBES {
					break // invalid game
				}
				inExpr = false
				red, blue, green = 0, 0, 0
				continue
			} else if char == '\n' {
				if red <= RED_CUBES && blue <= BLUE_CUBES && green <= GREEN_CUBES {
					res += gameCount // count valid game
				}
				break
			}

			if char >= '0' && char <= '9' {
				if !inExpr {
					inExpr = true
					currCountExpr = int(char - '0')
				} else {
					currCountExpr = currCountExpr*10 + int(char-'0')
				}
				continue
			}

			// if here, we are reading the color token
			if line[i+offset:i+offset+len("red")] == "red" {
				offset += 2 // only two because i is already incremented
				red = currCountExpr
			} else if line[i+offset:i+offset+len("blue")] == "blue" {
				offset += 3
				blue = currCountExpr
			} else if line[i+offset:i+offset+len("green")] == "green" {
				offset += 4
				green = currCountExpr
			}

			inExpr = false
		}
		gameCount++
	}

	return res
}

func second() int {
	fd, err := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("error opening input file -> %s\n", err)
		return 0
	}
	defer fd.Close()

	reader := bufio.NewReader(io.Reader(fd))

	res := 0
	for {
		lineB, err := reader.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0
		}

		line := strings.Split(string(lineB), ": ")[1]

		red, green, blue := 0, 0, 0
		offset := 0
		inExpr := false
		currExpr := 0
		for i := 0; i+offset < len(line); i++ {
			char := line[i+offset]
			if char == ' ' {
				continue
			} else if char == ',' {
				inExpr = false
				continue
			} else if char == ';' {
				inExpr = false
				continue
			} else if char == '\n' {
				res = res + red*green*blue
				break
			}

			if char >= '0' && char <= '9' {
				if !inExpr {
					inExpr = true
					currExpr = int(char - '0')
				} else {
					currExpr = 10*currExpr + int(char-'0')
				}
				continue
			}

			// if here, we are reading the color token
			if line[i+offset:i+offset+len("red")] == "red" {
				offset += 2 // only two because i is already incremented
				if currExpr > red {
					red = currExpr
				}
			} else if line[i+offset:i+offset+len("blue")] == "blue" {
				offset += 3
				if currExpr > blue {
					blue = currExpr
				}
			} else if line[i+offset:i+offset+len("green")] == "green" {
				offset += 4
				if currExpr > green {
					green = currExpr
				}
			}

			inExpr = false
		}
	}
	return res
}
