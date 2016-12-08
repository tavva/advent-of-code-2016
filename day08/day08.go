package main

import (
	"fmt"
	"io/ioutil"
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
		case "rotate":
		}
	}

	fmt.Println(screen)
}
