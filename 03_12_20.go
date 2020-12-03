package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("03_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice
	s := strings.Split(string(input), "\n")

	// Question #5
	numTrees := question5(s)
	fmt.Println(numTrees)

	// Question #6
	numTrees1 := question6(s, 1, 1)
	numTrees2 := question6(s, 3, 1)
	numTrees3 := question6(s, 5, 1)
	numTrees4 := question6(s, 7, 1)
	numTrees5 := question6(s, 1, 2)
	fmt.Println(numTrees1 * numTrees2 * numTrees3 * numTrees4 * numTrees5)
}

func question5(input []string) (numTrees int) {
	numTrees = 0
	trimmedInput := input[1:]
	index := 0
	for _, line := range trimmedInput {
		index += 3
		if index > len(line)-1 {
			index = index % (len(line))
		}
		if string(line[index]) == "#" {
			numTrees++
		}
	}
	return numTrees
}

func question6(input []string, right int, down int) (numTrees int) {
	numTrees = 0
	index := 0
	trimmedInput := input[1:]
	downCount := down
	for _, line := range trimmedInput {
		downCount--
		if downCount == 0 {
			downCount = down
		} else {
			continue
		}
		index += right
		if index > len(line)-1 {
			index = index % (len(line))
		}
		if string(line[index]) == "#" {
			numTrees++
		}
	}
	return numTrees
}
