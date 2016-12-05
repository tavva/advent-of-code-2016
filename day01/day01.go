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
	var current_location location

	var first_intercept location
	visited := make(map[location]bool)
	var matched = false

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

			current_location = location{x, y}

			if !matched && visited[current_location] {
				first_intercept = current_location
				matched = true
			}

			visited[location{x, y}] = true
		}
	}

	fmt.Printf("%v\n", math.Abs(x)+math.Abs(y))
	fmt.Printf("%v", math.Abs(first_intercept.x)+math.Abs(first_intercept.y))
}
