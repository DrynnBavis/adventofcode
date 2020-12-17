package main

import "fmt"

func main() {
	// Read input
	input := []int{10, 16, 6, 0, 1, 17}
	//input := []int{0, 3, 6}
	m := map[int]int{}
	for i, val := range input {
		m[val] = i
	}

	nextNum := 0
	thisNum := 0
	for count := len(input); count < 30000000; count++ {
		temp := 0
		if lastIndex, seen := m[nextNum]; seen {
			temp = count - lastIndex
		}
		//fmt.Println(count, nextNum, temp)
		m[nextNum] = count
		thisNum = nextNum
		nextNum = temp
	}
	fmt.Println(thisNum)
}
