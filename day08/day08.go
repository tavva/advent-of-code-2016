package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var screen [50][6]int

	input, _ := ioutil.ReadFile("./day08.txt")
	raw_lines := strings.TrimSpace(string(input))
	lines := strings.Split(raw_lines, "\n")

	for _, line := range lines {
	}
}
