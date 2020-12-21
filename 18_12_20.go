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
	input, err := ioutil.ReadFile("18_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice
	problems := strings.Split(string(input), "\n")

	// Solve problems
	totalSum := 0
	for _, problem := range problems {
		totalSum += solveProblemEqualPrec(problem)
	}
	fmt.Println(totalSum)

	totalSum = 0
	for _, problem := range problems {
		totalSum += solveProblemPlusPrec(problem)
	}
	fmt.Println(totalSum)
}

func solveProblemEqualPrec(problem string) (sum int) {
	operator := ""
	subProblem := ""
	writingToSub := false
	unclosed := 0
	for _, char := range problem {
		c := string(char)
		if c == " " {
			continue
		}
		if writingToSub {
			switch c {
			case "(":
				unclosed++
			case ")":
				if unclosed != 0 {
					unclosed--
				} else {
					result := solveProblemEqualPrec(subProblem)
					switch operator {
					case "+":
						sum += result
					case "*":
						sum *= result
					default:
						sum = result
					}
					subProblem = ""
					writingToSub = false
					continue
				}
			}
			if subProblem == "" {
				subProblem = c
			} else {
				subProblem += " "
				subProblem += c
			}
			continue
		}
		switch c {
		case "+":
			operator = c
		case "*":
			operator = c
		case "(":
			// start sub problem
			writingToSub = true
		default:
			val, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println("error in converting:", c)
			} else {
				switch operator {
				case "+":
					sum += val
				case "*":
					sum *= val
				default:
					sum = val
				}
			}
		}
	}
	return sum
}

func solveProblemPlusPrec(problem string) (sum int) {
	operator := ""
	subProblem := ""
	writingToSub := false
	unclosed := 0
	prevVal := 0
	multVals := []int{}
	temp := 0
	for _, char := range problem {
		c := string(char)
		if c == " " {
			continue
		}
		if writingToSub {
			switch c {
			case "(":
				unclosed++
			case ")":
				if unclosed != 0 {
					unclosed--
				} else {
					result := solveProblemPlusPrec(subProblem)
					prevVal = result
					subProblem = ""
					writingToSub = false
					continue
				}
			}
			if subProblem == "" {
				subProblem = c
			} else {
				subProblem += " "
				subProblem += c
			}
			continue
		}
		switch c {
		case "+":
			operator = c
			temp += prevVal
		case "*":
			operator = c
			if temp == 0 {
				multVals = append(multVals, prevVal)
			} else {
				temp += prevVal
				multVals = append(multVals, temp)
				temp = 0
			}
		case "(":
			// start sub problem
			writingToSub = true
		default:
			val, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println("error in converting:", c)
			} else {
				prevVal = val
			}
		}
	}
	switch operator {
	case "+":
		temp += prevVal
		multVals = append(multVals, temp)
	case "*":
		multVals = append(multVals, prevVal)
	}
	return getProduct(multVals)
}

func getProduct(vals []int) (product int) {
	product = 1
	for _, val := range vals {
		product *= val
	}
	return product
}
