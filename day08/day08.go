package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day08.txt")
	raw_lines := strings.TrimSpace(string(input))
	lines := strings.Split(raw_lines, "\n")

	for _, line := range lines {
	}
}
