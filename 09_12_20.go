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
	input, err := ioutil.ReadFile("09_12_20.txt")
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

	// Question #1:
	targetIndex := findOutlierIndex(nums)
	fmt.Println(nums[targetIndex])

	// Question #2:
	sumSlice := findContiguousSumSlice(nums, targetIndex)
	sort.Ints(sumSlice)
	fmt.Println(sumSlice[0] + sumSlice[len(sumSlice)-1])
}

func findOutlierIndex(nums []int) (outlier int) {
	for t := range nums {
		target := nums[t+25]
		found := false
		for left := t; left < t+24; left++ {
			for right := left + 1; right < t+25; right++ {
				if nums[left]+nums[right] == target {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return t + 25
		}
	}

	return 0
}

func findContiguousSumSlice(nums []int, targetIndex int) (sumSlice []int) {
	target := nums[targetIndex]
	for left := 0; left < targetIndex; left++ {
		count := 1
		for {
			right := left + count
			if sum(nums[left:right]) == target {
				return nums[left:right]
			}
			count++
			if right > targetIndex {
				break
			}
		}
	}

	return sumSlice
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
