package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func processTriangle(sides []int) {
	sort.Ints(sides)

	if sides[0]+sides[1] > sides[2] {
		fmt.Println(sides)
	}
}

func main() {
	var inputs [3][3]int

	var input, _ = ioutil.ReadFile("./day03.txt")
	var command_strings = strings.TrimSpace(string(input))
	var lines = strings.Split(command_strings, "\n")

	for row, line := range lines {
		var fields = strings.Fields(line)

		for col, v := range fields {
			var value, _ = strconv.Atoi(v)

			inputs[row%3][col] = value
		}

		if (row-2)%3 == 0 {
			for i := 0; i < 3; i++ {
				processTriangle([]int{
					inputs[0][i],
					inputs[1][i],
					inputs[2][i],
				})
			}
		}
	}
}
