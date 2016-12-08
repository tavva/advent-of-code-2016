package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func rect(screen [50][6]int, x int, y int) [50][6]int {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			screen[i][j] = 1
		}
	}

	return screen
}

func rotate_column(screen [50][6]int, coord int, amount int) [50][6]int {
	return screen
}
func rotate_row(screen [50][6]int, coord int, amount int) [50][6]int {
	return screen
}

func main() {
	var screen [50][6]int

	input, _ := ioutil.ReadFile("./day08.txt")
	raw_lines := strings.TrimSpace(string(input))
	lines := strings.Split(raw_lines, "\n")

	for _, line := range lines {
		tokens := strings.Fields(line)

		switch tokens[0] {
		case "rect":
			bits := strings.Split(tokens[1], "x")
			x, _ := strconv.Atoi(bits[0])
			y, _ := strconv.Atoi(bits[1])
			screen = rect(screen, x, y)
		case "rotate":
			amount, _ := strconv.Atoi(tokens[4])
			bits := strings.Split(tokens[2], "=")
			coord, _ := strconv.Atoi(bits[1])

			switch tokens[1] {
			case "column":
				screen = rotate_column(screen, coord, amount)
			case "row":
				screen = rotate_row(screen, coord, amount)
			}
		}
	}

	fmt.Println(screen)
}
