package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type coord struct {
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

	// Main Simulation Starts
	for cycle := 0; cycle < 6; cycle++ {
		activeWrites := []coord{}
		inactiveWrites := []coord{}

		// Pad the cube for extra space when adding new active cubes
		myCube = padMyCube(myCube)

		// Find all the changes we need to make
		for z, layer := range myCube {
			for y, row := range layer {
				for x, val := range row {
					// If a cube is active and exactly 2 or 3
					// of its neighbors are also active, the cube
					// remains active. Otherwise, the cube becomes
					// inactive.
					// If a cube is inactive but exactly 3 of its
					// neighbors are active, the cube becomes active.
					// Otherwise, the cube remains inactive.
					myCoord := coord{z, y, x}
					numActive := getActiveNeighbours(myCube, myCoord)
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

		// Apply changes
		for _, activeCoord := range activeWrites {
			myCube[activeCoord.z][activeCoord.y][activeCoord.x] = "#"
		}
		for _, inactiveCoord := range inactiveWrites {
			myCube[inactiveCoord.z][inactiveCoord.y][inactiveCoord.x] = "."
		}
	}

	// Print num active cells
	fmt.Println(countActive(myCube))
}

func getActiveNeighbours(cube [][][]string, pos coord) (numActive int) {
	limitZ := len(cube) - 1
	limitY := len(cube[0]) - 1
	limitX := len(cube[0][0]) - 1

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

	for _, z := range zSpace {

		for _, y := range ySpace {

			for _, x := range xSpace {
				if cube[z][y][x] == "#" && (coord{z, y, x} != pos) {
					numActive++
				}
			}
		}
	}

	return numActive
}

func padMyCube(cube [][][]string) (newCube [][][]string) {
	padZBot := false
	padZTop := false
	padYBot := false
	padYTop := false
	padXLeft := false
	padXRight := false
	for z, layer := range cube {
		for y, row := range layer {
			for x, val := range row {
				if val == "#" {
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

	for z := range cube {
		for y := range cube[z] {
			if padXLeft {
				cube[z][y] = append([]string{"."}, cube[z][y]...)
			}
			if padXRight {
				cube[z][y] = append(cube[z][y], ".")
			}
		}
		if padYBot {
			newRow := []string{}
			for i := 0; i < len(cube[z][0]); i++ {
				newRow = append(newRow, ".")
			}
			cube[z] = append([][]string{newRow}, cube[z]...)
		}
		if padYTop {
			newRow := []string{}
			for i := 0; i < len(cube[z][0]); i++ {
				newRow = append(newRow, ".")
			}
			cube[z] = append(cube[z], newRow)
		}
	}
	if padZBot {
		newLayer := [][]string{}
		for k := 0; k < len(cube[0]); k++ {
			newRow := []string{}
			for i := 0; i < len(cube[0][0]); i++ {
				newRow = append(newRow, ".")
			}
			newLayer = append(newLayer, newRow)
		}
		cube = append([][][]string{newLayer}, cube...)
	}
	if padZTop {
		newLayer := [][]string{}
		for k := 0; k < len(cube[0]); k++ {
			newRow := []string{}
			for i := 0; i < len(cube[0][0]); i++ {
				newRow = append(newRow, ".")
			}
			newLayer = append(newLayer, newRow)
		}
		cube = append(cube, newLayer)
	}
	return cube
}

func printCube(myCube [][][]string) {
	for _, layer := range myCube {
		for _, row := range layer {
			fmt.Println(row)
		}
		fmt.Println("------")
	}
}

func countActive(myCube [][][]string) (total int) {
	for _, layer := range myCube {
		for _, row := range layer {
			for _, val := range row {
				if val == "#" {
					total++
				}
			}
		}
	}
	return total
}
