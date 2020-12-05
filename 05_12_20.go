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
	input, err := ioutil.ReadFile("05_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")

	// Question #1
	highestSeatID := 0
	for _, line := range s {
		seatID := getSeatID(line)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	fmt.Println(highestSeatID)

	// Question #2
	ids := []int{}
	for _, line := range s {
		seatID := getSeatID(line)
		ids = append(ids, seatID)
	}
	sort.Ints(ids)
	prevID := 0
	for _, seatID := range ids {
		if prevID != 0 {
			if seatID != prevID+1 {
				fmt.Println(seatID - 1)
			}
		}
		prevID = seatID
	}

}

func getSeatID(input string) (seatID int) {
	rowPath := input[:7]
	colPath := input[7:]
	startingRow := ""
	startingCol := ""

	for _, c := range rowPath {
		if string(c) == "F" {
			startingRow += "0"
		} else {
			startingRow += "1"
		}
	}
	rowNumber, err := strconv.ParseInt(startingRow, 2, 32)
	if err != nil {
		fmt.Println("Failed to parse ", startingRow)
	}

	for _, c := range colPath {
		if string(c) == "L" {
			startingCol += "0"
		} else {
			startingCol += "1"
		}
	}
	colNumber, err := strconv.ParseInt(startingCol, 2, 32)
	if err != nil {
		fmt.Println("Failed to parse ", startingCol)
	}
	return int(8*rowNumber + colNumber)
}
