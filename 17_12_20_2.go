package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type coord struct {
	w int
	z int
	y int
	x int
}

func main() {
	// Read input
	input, err := ioutil.ReadFile("17_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}
	myHyper := [][][][]string{}
	myCube := [][][]string{}

	// Parse into slice
	s := strings.Split(string(input), "\n")
	layer := [][]string{}
	for _, line := range s {
		row := []string{}
		for _, c := range line {
			row = append(row, string(c))
		}
		layer = append(layer, row)
	}
	myCube = append(myCube, layer)
	myHyper = append(myHyper, myCube)

	// Main Simulation Starts
	for cycle := 0; cycle < 6; cycle++ {
		activeWrites := []coord{}
		inactiveWrites := []coord{}

		// Pad the cube for extra space when adding new active cubes
		myHyper = padMyHyper(myHyper)

		// Find all the changes we need to make
		for w, cube := range myHyper {
			for z, layer := range cube {
				for y, row := range layer {
					for x, val := range row {
						// If a cube is active and exactly 2 or 3
						// of its neighbors are also active, the cube
						// remains active. Otherwise, the cube becomes
						// inactive.
						// If a cube is inactive but exactly 3 of its
						// neighbors are active, the cube becomes active.
						// Otherwise, the cube remains inactive.
						myCoord := coord{w, z, y, x}
						numActive := getActiveNeighbours(myHyper, myCoord)
						if val == "." {
							// inactive case
							if numActive == 3 {
								activeWrites = append(activeWrites, myCoord)
							}
						} else {
							// active case
							if numActive < 2 || numActive > 3 {
								inactiveWrites = append(inactiveWrites, myCoord)
							}
						}
					}
				}
			}
		}

		// Apply changes
		for _, activeCoord := range activeWrites {
			myHyper[activeCoord.w][activeCoord.z][activeCoord.y][activeCoord.x] = "#"
		}
		for _, inactiveCoord := range inactiveWrites {
			myHyper[inactiveCoord.w][inactiveCoord.z][inactiveCoord.y][inactiveCoord.x] = "."
		}
	}

	// Print num active cells
	fmt.Println(countActive(myHyper))
}

func getActiveNeighbours(hyper [][][][]string, pos coord) (numActive int) {
	limitW := len(hyper) - 1
	limitZ := len(hyper[0]) - 1
	limitY := len(hyper[0][0]) - 1
	limitX := len(hyper[0][0][0]) - 1

	// Get W iteration space
	wSpace := []int{}
	if pos.w != 0 {
		wSpace = append(wSpace, pos.w-1)
	}
	wSpace = append(wSpace, pos.w)
	if pos.w != limitW {
		wSpace = append(wSpace, pos.w+1)
	}

	// Get Z iteration space
	zSpace := []int{}
	if pos.z != 0 {
		zSpace = append(zSpace, pos.z-1)
	}
	zSpace = append(zSpace, pos.z)
	if pos.z != limitZ {
		zSpace = append(zSpace, pos.z+1)
	}

	// Get Y iteration space
	ySpace := []int{}
	if pos.y != 0 {
		ySpace = append(ySpace, pos.y-1)
	}
	ySpace = append(ySpace, pos.y)
	if pos.y != limitY {
		ySpace = append(ySpace, pos.y+1)
	}

	// Get X iteration space
	xSpace := []int{}
	if pos.x != 0 {
		xSpace = append(xSpace, pos.x-1)
	}
	xSpace = append(xSpace, pos.x)
	if pos.x != limitX {
		xSpace = append(xSpace, pos.x+1)
	}

	for _, w := range wSpace {
		for _, z := range zSpace {
			for _, y := range ySpace {
				for _, x := range xSpace {
					if hyper[w][z][y][x] == "#" && (coord{w, z, y, x} != pos) {
						numActive++
					}
				}
			}
		}
	}

	return numActive
}

func padMyHyper(hyper [][][][]string) (newHyper [][][][]string) {
	padWBot := false
	padWTop := false
	padZBot := false
	padZTop := false
	padYBot := false
	padYTop := false
	padXLeft := false
	padXRight := false
	for w, cube := range hyper {
		for z, layer := range cube {
			for y, row := range layer {
				for x, val := range row {
					if val == "#" {
						if w == 0 {
							padWBot = true
						}
						if w == len(hyper)-1 {
							padWTop = true
						}
						if z == 0 {
							padZBot = true
						}
						if z == len(cube)-1 {
							padZTop = true
						}
						if y == 0 {
							padYBot = true
						}
						if y == len(layer)-1 {
							padYTop = true
						}
						if x == 0 {
							padXLeft = true
						}
						if x == len(row)-1 {
							padXRight = true
						}
					}
				}
			}
		}
	}

	for w := range hyper {
		for z := range hyper[w] {
			for y := range hyper[w][z] {
				if padXLeft {
					hyper[w][z][y] = append([]string{"."}, hyper[w][z][y]...)
				}
				if padXRight {
					hyper[w][z][y] = append(hyper[w][z][y], ".")
				}
			}
			if padYBot {
				newRow := []string{}
				for i := 0; i < len(hyper[w][z][0]); i++ {
					newRow = append(newRow, ".")
				}
				hyper[w][z] = append([][]string{newRow}, hyper[w][z]...)
			}
			if padYTop {
				newRow := []string{}
				for i := 0; i < len(hyper[w][z][0]); i++ {
					newRow = append(newRow, ".")
				}
				hyper[w][z] = append(hyper[w][z], newRow)
			}
		}
		if padZBot {
			newLayer := [][]string{}
			for k := 0; k < len(hyper[w][0]); k++ {
				newRow := []string{}
				for i := 0; i < len(hyper[w][0][0]); i++ {
					newRow = append(newRow, ".")
				}
				newLayer = append(newLayer, newRow)
			}
			hyper[w] = append([][][]string{newLayer}, hyper[w]...)
		}
		if padZTop {
			newLayer := [][]string{}
			for k := 0; k < len(hyper[w][0]); k++ {
				newRow := []string{}
				for i := 0; i < len(hyper[w][0][0]); i++ {
					newRow = append(newRow, ".")
				}
				newLayer = append(newLayer, newRow)
			}
			hyper[w] = append(hyper[w], newLayer)
		}
	}
	if padWBot {
		newCube := [][][]string{}
		for z := 0; z < len(hyper[0]); z++ {
			newLayer := [][]string{}
			for y := 0; y < len(hyper[0][0]); y++ {
				newRow := []string{}
				for x := 0; x < len(hyper[0][0][0]); x++ {
					newRow = append(newRow, ".")
				}
				newLayer = append(newLayer, newRow)
			}
			newCube = append(newCube, newLayer)
		}
		hyper = append([][][][]string{newCube}, hyper...)
	}
	if padWTop {
		newCube := [][][]string{}
		for z := 0; z < len(hyper[0]); z++ {
			newLayer := [][]string{}
			for y := 0; y < len(hyper[0][0]); y++ {
				newRow := []string{}
				for x := 0; x < len(hyper[0][0][0]); x++ {
					newRow = append(newRow, ".")
				}
				newLayer = append(newLayer, newRow)
			}
			newCube = append(newCube, newLayer)
		}
		hyper = append(hyper, newCube)
	}

	return hyper
}

func printCube(hyper [][][][]string) {
	for _, cube := range hyper {
		for _, layer := range cube {
			for _, row := range layer {
				fmt.Println(row)
			}
			fmt.Println("------")
		}
	}
}

func countActive(hyper [][][][]string) (total int) {
	for _, cube := range hyper {
		for _, layer := range cube {
			for _, row := range layer {
				for _, val := range row {
					if val == "#" {
						total++
					}
				}
			}
		}
	}
	return total
}
