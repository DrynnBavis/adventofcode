package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("01_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into sorted slice
	s := strings.Split(string(input), "\n")
	nums := []int{}
	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, j)
	}
	sort.Ints(nums)

	// Question #1
	num1, num2 := question1(nums, 2020)
	fmt.Println(num1 * num2)

	// Question #2
	num1, num2, num3 := question2(nums)
	fmt.Println(num1 * num2 * num3)
}

func question1(nums []int, target int) (int, int) {
	for i := range nums {
		front := nums[i]
		for k := len(nums) - 1; k >= 0; k-- {
			back := nums[k]
			if front+back < target {
				break
			}
			if front+back == target {
				return front, back
			}
		}
	}
	return 0, 0
}

func question2(nums []int) (int, int, int) {
	for i := range nums {
		remainder := 2020 - nums[i]
		num1, num2 := question1(nums[i:], remainder)
		if num1 != 0 && num2 != 0 {
			return nums[i], num1, num2
		}
	}
	return 0, 0, 0
}
