package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
			for i := 0; i < x; i++ {
				for j := 0; j < y; j++ {
					screen[i][j] = 1
				}
			}
		case "rotate":
		}
	}

	fmt.Println(screen)
}
