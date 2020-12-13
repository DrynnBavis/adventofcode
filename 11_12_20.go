package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("11_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")
	parsed1 := [][]string{}
	parsed2 := [][]string{}
	for _, line := range s {
		newLine := []string{}
		for _, c := range line {
			newLine = append(newLine, string(c))
		}
		parsed1 = append(parsed1, newLine)
		parsed2 = append(parsed2, newLine)
	}

	// Question #1:
	// stable := stabiliseSeatChaos1(parsed1)
	// filledCount := 0
	// for _, line := range stable {
	// 	for _, c := range line {
	// 		if string(c) == "#" {
	// 			filledCount++
	// 		}
	// 	}
	// }
	// fmt.Println(filledCount)

	// Quenstion #2:
	stable := stabiliseSeatChaos2(parsed2)
	filledCount := 0
	for _, line := range stable {
		for _, c := range line {
			if string(c) == "#" {
				filledCount++
			}
		}
	}
	fmt.Println(filledCount)
}

type coord struct {
	y int
	x int
}

func stabiliseSeatChaos1(input [][]string) (stableSeats [][]string) {
	LIM_LEFT := 0
	LIM_TOP := 0
	LIM_RIGHT := len(input[0]) - 1
	LIM_BOTTOM := len(input) - 1
	emptyWrites := []coord{}
	fillWrites := []coord{}

	for {
		changes := 0
		for y := LIM_TOP; y <= LIM_BOTTOM; y++ {
			for x := LIM_LEFT; x <= LIM_RIGHT; x++ {
				sym := string(input[y][x])
				switch sym {
				case "L":
					adjacentEmptyCount := 0
					for j := y - 1; j <= y+1; j++ {
						for k := x - 1; k <= x+1; k++ {
							if j < LIM_TOP || j > LIM_BOTTOM || k < LIM_LEFT || k > LIM_RIGHT {
								adjacentEmptyCount++
							} else {
								if j >= LIM_TOP && j <= LIM_BOTTOM && k >= LIM_LEFT && k <= LIM_RIGHT && !(j == y && k == x) {
									sideSeat := string(input[j][k])
									if sideSeat == "L" || sideSeat == "." {
										adjacentEmptyCount++
									}
								}
							}
						}
					}
					if adjacentEmptyCount == 8 {
						fillWrites = append(fillWrites, coord{y, x})
						changes++
					}
				case "#":
					adjacentEmptyCount := 0
					for j := y - 1; j <= y+1; j++ {
						for k := x - 1; k <= x+1; k++ {
							if j < LIM_TOP || j > LIM_BOTTOM || k < LIM_LEFT || k > LIM_RIGHT {
								adjacentEmptyCount++
							} else {
								if j >= LIM_TOP && j <= LIM_BOTTOM && k >= LIM_LEFT && k <= LIM_RIGHT && !(j == y && k == x) {
									sideSeat := string(input[j][k])
									if sideSeat == "L" || sideSeat == "." {
										adjacentEmptyCount++
									}
								}
							}
						}
					}
					if adjacentEmptyCount <= 4 {
						emptyWrites = append(emptyWrites, coord{y, x})
						changes++
					}
				default:
					//floor
				}
			}
		}
		// fmt.Println("empty", len(emptyWrites))
		// fmt.Println("fill", len(fillWrites))
		// fmt.Println(changes)
		// for _, line := range input {
		// 	fmt.Println(line)
		// }
		// apply writes
		for _, coord := range emptyWrites {
			input[coord.y][coord.x] = "L"
		}
		for _, coord := range fillWrites {
			input[coord.y][coord.x] = "#"
		}

		// fmt.Println("------------------------")
		// for _, line := range input {
		// 	fmt.Println(line)
		// }
		if changes == 0 {
			break
		}
		emptyWrites = nil
		fillWrites = nil
	}
	return input
}

func stabiliseSeatChaos2(input [][]string) (stableSeats [][]string) {
	LIM_LEFT := 0
	LIM_TOP := 0
	LIM_RIGHT := len(input[0]) - 1
	LIM_BOTTOM := len(input) - 1
	emptyWrites := []coord{}
	fillWrites := []coord{}

	for {
		changes := 0
		for y := LIM_TOP; y <= LIM_BOTTOM; y++ {
			for x := LIM_LEFT; x <= LIM_RIGHT; x++ {
				sym := string(input[y][x])
				if sym == "." {
					continue
				}
				numEmpty := 0
				// TOP
				j := 1
				for {
					if y-j < LIM_TOP {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y-j][x] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// BOT
				j = 1
				for {
					if y+j > LIM_BOTTOM {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y+j][x] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// LEFT
				j = 1
				for {
					if x-j < LIM_LEFT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y][x-j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// RIGHT
				j = 1
				for {
					if x+j > LIM_RIGHT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y][x+j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// TOP RIGHT
				j = 1
				for {
					if y-j < LIM_TOP {
						numEmpty++
						break
					}
					if x+j > LIM_RIGHT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y-j][x+j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// BOT RIGHT
				j = 1
				for {
					if y+j > LIM_BOTTOM {
						numEmpty++
						break
					}
					if x+j > LIM_RIGHT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y+j][x+j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// TOP LEFT
				j = 1
				for {
					if y-j < LIM_TOP {
						numEmpty++
						break
					}
					if x-j < LIM_LEFT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y-j][x-j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				// BOT LEFT
				j = 1
				for {
					if y+j > LIM_BOTTOM {
						numEmpty++
						break
					}
					if x-j < LIM_LEFT {
						numEmpty++
						break
					}
					foundSomething := false
					switch input[y+j][x-j] {
					case "#":
						foundSomething = true
					case "L":
						numEmpty++
						foundSomething = true
					}
					if foundSomething {
						break
					}
					j++
				}
				if sym == "L" {
					if numEmpty == 8 {
						fillWrites = append(fillWrites, coord{y, x})
						changes++
					}
				} else {
					if numEmpty <= 3 {
						emptyWrites = append(emptyWrites, coord{y, x})
						changes++
					}
				}
			}
		}

		// apply writes
		for _, coord := range emptyWrites {
			input[coord.y][coord.x] = "L"
		}
		for _, coord := range fillWrites {
			input[coord.y][coord.x] = "#"
		}

		// stop when we stabilise
		if changes == 0 {
			break
		}
		emptyWrites = nil
		fillWrites = nil
	}
	return input
}
