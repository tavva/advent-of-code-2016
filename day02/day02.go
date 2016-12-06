package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day02.txt")
	var lines = strings.Split(string(input), "\n")

	for _, line := range lines {
		fmt.Printf("%v", line)
	}
}