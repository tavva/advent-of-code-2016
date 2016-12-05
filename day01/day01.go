package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type location struct {
	x, y float64
}

func main() {
	var input, _ = ioutil.ReadFile("./day01.txt")
	var command_string = strings.TrimSpace(string(input))

	var commands = strings.Split(command_string, ", ")

	var x, y float64 // defaults to 0
	var current_direction int

	visited := make(map[location]bool)

	for _, v := range commands {
		var direction = string(v[0])
		var distance, _ = strconv.Atoi(string(v[1:len(v)]))

		if direction == "R" {
			current_direction++
		} else if direction == "L" {
			current_direction += -1 + 4
		}

		current_direction %= 4

		for i := 0; i < distance; i++ {
			switch current_direction {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			}

			visited[location{x, y}] = true
		}
	}

	fmt.Printf("%v", math.Abs(x)+math.Abs(y))
}
