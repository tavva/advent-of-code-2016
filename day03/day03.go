package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day03.txt")
	var command_strings = strings.TrimSpace(string(input))
	var lines = strings.Split(command_strings, "\n")

	for _, line := range lines {
		var side_inputs = strings.Fields(line)
		var sides []int

		for _, v := range side_inputs {
			var value, _ = strconv.Atoi(v)
			sides = append(sides, value)
		}

		sort.Ints(sides)
		fmt.Println(sides)
	}
}
