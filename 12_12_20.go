package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type move struct {
	dir string
	val int
}

type coord struct {
	x int
	y int
}

func main() {
	// Read input
	input, err := ioutil.ReadFile("12_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")
	moves := []move{}
	for _, line := range s {
		dir := string(line[0])
		val, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if err != nil {
			fmt.Println("Failed to parse ", line[1:])
		}
		moves = append(moves, move{dir, int(val)})
	}

	// Question #1:
	fmt.Println(getManhattanDist(moves))

	// Question #2:
	fmt.Println(getManhattanDistByWaypoint(moves))
}

func getManhattanDist(input []move) (dist int) {
	longitude := 0 // +North,-South
	latitude := 0  // +East/-West
	facing := 0
	for _, move := range input {
		switch move.dir {
		case "N":
			longitude += move.val
		case "S":
			longitude -= move.val
		case "E":
			latitude += move.val
		case "W":
			latitude -= move.val
		case "F":
			switch intToDir(facing) {
			case "N":
				longitude += move.val
			case "S":
				longitude -= move.val
			case "E":
				latitude += move.val
			case "W":
				latitude -= move.val
			}
		case "L":
			facing += move.val / 90
		case "R":
			facing -= move.val / 90
		}
	}
	dist = Abs(longitude) + Abs(latitude)
	return dist
}

func getManhattanDistByWaypoint(input []move) (dist int) {
	shipPos := coord{0, 0}
	waypoint := coord{10, 1}

	for _, move := range input {
		switch move.dir {
		case "N":
			waypoint.y += move.val
		case "S":
			waypoint.y -= move.val
		case "E":
			waypoint.x += move.val
		case "W":
			waypoint.x -= move.val
		case "F":
			xGap := move.val * (waypoint.x)
			yGap := move.val * (waypoint.y)
			shipPos.x += xGap
			shipPos.y += yGap
		case "L":
			waypoint.x, waypoint.y = Rotate(waypoint.x, waypoint.y, move.val)
		case "R":
			waypoint.x, waypoint.y = Rotate(waypoint.x, waypoint.y, move.val*-1)
		}
	}
	dist = Abs(shipPos.x) + Abs(shipPos.y)
	return dist
}

func intToDir(face int) (dir string) {
	if face < 0 {
		face *= -1
		face += 2
	}
	switch face % 4 {
	case 0:
		dir = "E"
	case 1:
		dir = "N"
	case 2:
		dir = "W"
	case 3:
		dir = "S"
	}
	return dir
}

// Abs ...
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Rotate ...
func Rotate(x0 int, y0 int, theta int) (x1 int, y1 int) {
	rad := (float64(theta) * (math.Pi / 180.0))
	x1 = int(math.Cos(rad))*x0 - int(math.Sin(rad))*y0
	y1 = int(math.Sin(rad))*x0 + int(math.Cos(rad))*y0
	return int(x1), int(y1)
}
