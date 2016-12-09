package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day09.txt")
	data := strings.TrimSpace(string(input))

	for i := 0; i < len(data); i++ {
		r := rune(data[i])

		switch r {
		default:
			fmt.Printf("%v", string(r))
		}
	}
}
