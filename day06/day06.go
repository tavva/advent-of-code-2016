package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day06-short.txt")
	var raw_lines = strings.TrimSpace(string(input))
	var lines = strings.Split(raw_lines, "\n")

	for _, line := range lines {
		println(line)
	}
}
