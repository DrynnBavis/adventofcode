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
	input, err := ioutil.ReadFile("10_12_20.txt")
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

	// Question #1:
	gaps := (countGaps(nums))
	fmt.Println(gaps[1] * gaps[3])

	// Question #2"
	nums = append([]int{0}, nums...)
	nums = append(nums, nums[len(nums)-1]+3)
	fmt.Println(countStackCombos(nums))

}

func countGaps(adapters []int) (counts map[int]int) {
	prevRating := 0
	counts = map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}
	for _, adapterRating := range adapters {
		gap := adapterRating - prevRating
		prevRating = adapterRating
		counts[gap]++
	}
	return counts
}

func countStackCombos(adapters []int) (count int) {
	numCombos := map[int]int{}
	// find our forks and mark them in our empty map
	for i := range adapters {
		if k := i + 2; k < len(adapters) {
			gap := adapters[k] - adapters[i]
			if gap <= 3 {
				numCombos[adapters[i]] = 0
			}
		}
	}

	// we know there's only one path at the end of list!
	end := len(adapters) - 1
	numCombos[adapters[end]] = 1

	// back track and memoie the combos at each index
	for i := end - 1; i >= 0; i-- {
		if _, ok := numCombos[adapters[i]]; ok {
			// we found a fork
			for k := 0; k <= 3 && i+k <= end; k++ {
				gap := adapters[i+k] - adapters[i]
				if gap <= 3 {
					numCombos[adapters[i]] += numCombos[adapters[i+k]]
				}
			}
		} else {
			numCombos[adapters[i]] = numCombos[adapters[i+1]]
		}
	}
	return numCombos[0]
}
