package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var input, _ = ioutil.ReadFile("./day01.txt")
	var command_string = strings.TrimSpace(string(input))

	var commands = strings.Split(command_string, ", ")

	fmt.Printf("%v", commands)
}
