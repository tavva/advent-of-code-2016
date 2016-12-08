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
	for i := 0; i < amount; i++ {
		var temp_col [6]int

		for j := 0; j < 6; j++ {
			temp_col[(j+1)%6] = screen[coord][j]
		}
		for j := 0; j < 6; j++ {
			screen[coord][j] = temp_col[j]
		}
	}

	return screen
}

func rotate_row(screen [50][6]int, coord int, amount int) [50][6]int {
	for i := 0; i < amount; i++ {
		var temp_row [50]int

		for j := 0; j < 50; j++ {
			temp_row[(j+1)%50] = screen[j][coord]
		}
		for j := 0; j < 50; j++ {
			screen[j][coord] = temp_row[j]
		}
	}

	return screen
}

func print_screen(screen [50][6]int) {
	for i := 0; i < len(screen[0]); i++ {
		for j := 0; j < len(screen); j++ {
			if j%5 == 0 {
				fmt.Printf(" ")
			}
			if screen[j][i] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%v", screen[j][i])
			}
		}
		fmt.Printf("\n")
	}
}

func count_lit(screen [50][6]int) int {
	num_lit := 0

	for i := 0; i < len(screen[0]); i++ {
		for j := 0; j < len(screen); j++ {
			if screen[j][i] == 1 {
				num_lit++
			}
		}
	}

	return num_lit
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

	print_screen(screen)
	println(count_lit(screen))
}
