package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day03.txt")
	var command_strings = strings.TrimSpace(string(input))
	var lines = strings.Split(command_strings, "\n")

	for _, line := range lines {
		println(line)
	}
}
