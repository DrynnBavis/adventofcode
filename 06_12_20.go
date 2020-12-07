package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("06_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse input into 1 string per group
	s := strings.Split(string(input), "\n")
	groups := []string{}
	newLine := true
	count := 0
	for _, line := range s {
		if line == "" {
			newLine = true
			count++
		} else {
			if newLine {
				groups = append(groups, line)
			} else {
				groups[count] += " " + line
			}
			newLine = false
		}
	}

	// Question #1
	totalAnyYes := 0
	totalAllYes := 0
	for _, group := range groups {
		anyYes, allYes := getAnyAllYesFromGroup(group)
		totalAnyYes += anyYes
		totalAllYes += allYes
	}
	fmt.Println(totalAnyYes)
	fmt.Println(totalAllYes)
}

func getAnyAllYesFromGroup(input string) (anyYes int, allYes int) {
	m := map[string]int{}
	people := strings.Split(string(input), " ")
	for _, person := range people {
		for _, c := range person {
			key := string(c)
			if val, ok := m[key]; ok {
				m[key] = val + 1
			} else {
				m[key] = 1
			}
		}
	}
	anyYes = len(m)
	for _, val := range m {
		if val == len(people) {
			allYes++
		}
	}
	return anyYes, allYes
}
