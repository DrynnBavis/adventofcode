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
	input, err := ioutil.ReadFile("07_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse input into 1 string per group
	s := strings.Split(string(input), "\n")

	// Question #1
	fmt.Println(getNumOuterBags(s, "shiny gold"))

	// Question #2
	fmt.Println(getNumInnerBags(s, "shiny gold"))
}

func getNumOuterBags(input []string, rootBag string) (numOuterBags int) {
	bags := map[string](map[string]int){}
	for _, line := range input {
		// get master bag name
		parentColour := strings.Split(string(line), " bags contain ")[0]
		remainder := strings.Split(string(line), " bags contain ")[1]
		childBags := strings.Split(string(remainder), ", ")
		// get children
		for _, child := range childBags {
			if string(child) == "no other bags." {
				continue
			}
			capacity, err := strconv.ParseInt(string(child[0]), 10, 32)
			if err != nil {
				fmt.Println("Failed to parse ", string(child[0]))
			}
			words := strings.Split(string(child), " ")
			childColour := string(words[1]) + " " + string(words[2])
			if _, ok := bags[childColour]; !ok {
				bags[childColour] = map[string]int{}
			}
			bags[childColour][parentColour] = int(capacity)

		}
	}
	parentBags := map[string]int{}
	bagsToCheck := []string{}
	for key, val := range bags[rootBag] {
		bagsToCheck = append(bagsToCheck, key)
		parentBags[key] = val
	}
	for {
		if len(bagsToCheck) == 0 {
			break
		}
		newBag := bagsToCheck[0]
		bagsToCheck[0] = ""
		bagsToCheck = bagsToCheck[1:]
		for key, val := range bags[newBag] {
			bagsToCheck = append(bagsToCheck, key)
			parentBags[key] = val
		}
	}
	return len(parentBags)
}

func getNumInnerBags(input []string, rootBag string) (numInnerBags int) {
	bags := map[string](map[string]int){}
	for _, line := range input {
		// get master bag name
		parentColour := strings.Split(string(line), " bags contain ")[0]
		remainder := strings.Split(string(line), " bags contain ")[1]
		childBags := strings.Split(string(remainder), ", ")
		// get children
		if string(childBags[0]) == "no other bags." {
			continue
		}
		m := map[string]int{}
		for _, child := range childBags {
			capacity, err := strconv.ParseInt(string(child[0]), 10, 32)
			if err != nil {
				fmt.Println("Failed to parse ", string(child[0]))
			}
			words := strings.Split(string(child), " ")
			childColour := string(words[1]) + " " + string(words[2])
			m[childColour] = int(capacity)
		}
		bags[parentColour] = m
	}
	numInnerBags = helpCountInner(bags, rootBag)
	return numInnerBags - 1
}

func helpCountInner(bags map[string](map[string]int), target string) (count int) {
	count = 1
	for bagName, bagCount := range bags[target] {
		count += bagCount * helpCountInner(bags, bagName)
	}
	return count
}
