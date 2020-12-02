package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("02_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice
	s := strings.Split(string(input), "\n")

	// Question #3
	correctPasswords := question3(s)
	fmt.Println(len(correctPasswords))

	// Question #4
	correctPasswords = question4(s)
	fmt.Println(len(correctPasswords))
}

func question3(input []string) (correctPasswords []string) {
	results := []string{}
	for _, line := range input {
		l := strings.Split(line, " ")
		limits := strings.Split(l[0], "-")
		minCount, err := strconv.Atoi(limits[0])
		if err != nil {
			panic(err)
		}
		maxCount, err := strconv.Atoi(limits[1])
		if err != nil {
			panic(err)
		}
		key := l[1]
		pattern := l[2]
		hits := 0
		for _, c := range pattern {
			if byte(c) == key[0] {
				hits++
			}
		}
		if hits >= minCount && hits <= maxCount {
			results = append(results, pattern)
		}
	}
	return results
}

func question4(input []string) (correctPasswords []string) {
	results := []string{}
	for _, line := range input {
		l := strings.Split(line, " ")
		limits := strings.Split(l[0], "-")
		firstIndex, err := strconv.Atoi(limits[0])
		if err != nil {
			panic(err)
		}
		secondIndex, err := strconv.Atoi(limits[1])
		if err != nil {
			panic(err)
		}
		key := l[1]
		pattern := l[2]
		if (pattern[firstIndex-1] == key[0]) != (pattern[secondIndex-1] == key[0]) {
			results = append(results, pattern)
		}
	}
	return results
}
