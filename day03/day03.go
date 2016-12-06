package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day03.txt")
	var command_strings = strings.TrimSpace(string(input))
	var lines = strings.Split(command_strings, "\n")

	for _, line := range lines {
		var side_inputs = strings.Fields(line)
		var sides [3]int

		for x, v := range side_inputs {
			sides[x], _ = strconv.Atoi(v)
		}
	}
}
