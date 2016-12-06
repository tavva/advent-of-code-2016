package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	var pin string

	keypad := [5][5]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	var x, y = 0, 2 // coordinates on keypad
	var temp_x, temp_y int

	var input, _ = ioutil.ReadFile("./day02.txt")
	var command_string = strings.TrimSpace(string(input))
	var lines = strings.Split(command_string, "\n")

	for _, line := range lines {
		for _, direction := range line {
			temp_x = x
			temp_y = y

			switch direction {
			case 'U':
				temp_y = int(math.Max(0, float64(y)-1))
			case 'D':
				temp_y = int(math.Min(4, float64(y)+1))
			case 'L':
				temp_x = int(math.Max(0, float64(x)-1))
			case 'R':
				temp_x = int(math.Min(4, float64(x)+1))
			}

			if keypad[temp_y][temp_x] != "" {
				x = temp_x
				y = temp_y
			}
		}

		pin = pin + keypad[y][x]
	}

	println(pin)
}
