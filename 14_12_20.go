package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	input, err := ioutil.ReadFile("14_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")

	// Question #1:
	fmt.Println(getSumOfMem(s))

	// Question #2:
	fmt.Println(getSumOfMem2(s))
}

func getSumOfMem(input []string) (sum int64) {
	mem := map[int]int64{}
	mask := ""
	for _, line := range input {
		l := strings.Split(line, " = ")
		switch l[0] {
		case "mask":
			// update mask
			mask = l[1]
		default:
			// write mem
			addrStr := l[0][4 : len(l[0])-1]
			addr, err := strconv.Atoi(addrStr)
			if err != nil {
				fmt.Println("Failed to parse", addr)
			}
			val, err := strconv.Atoi(l[1])
			if err != nil {
				fmt.Println("Failed to parse", l[1])
			}
			maskedVal := maskInt32(val, mask)

			mem[addr] = maskedVal
		}
	}

	for _, val := range mem {
		sum += val
	}

	return sum
}
func getSumOfMem2(input []string) (sum int64) {
	mem := map[int64]int64{}
	mask := ""
	for _, line := range input {
		l := strings.Split(line, " = ")
		switch l[0] {
		case "mask":
			// update mask
			mask = l[1]
		default:
			// write mem
			addrStr := l[0][4 : len(l[0])-1]
			addr, err := strconv.ParseInt(addrStr, 10, 36)
			if err != nil {
				fmt.Println("Failed to parse", addr)
			}
			addresses := mask2Int32(addr, mask)

			val, err := strconv.ParseInt(l[1], 10, 36)
			if err != nil {
				fmt.Println("Failed to parse", l[1])
			}

			for _, addr := range addresses {
				mem[addr] = val
			}
		}
	}

	for _, val := range mem {
		sum += val
	}

	return sum
}

func maskInt32(source int, mask string) (maskedInt int64) {
	result := ""
	input := strconv.FormatInt(int64(source), 2)
	input2 := ""
	for i := len(input); i < 36; i++ {
		input2 += "0"
	}
	input2 = input2 + input
	for i, c := range mask {
		switch string(c) {
		case "1":
			result += "1"
		case "0":
			result += "0"
		default:
			result += string(input2[i])
		}
	}
	maskedVal, err := strconv.ParseInt(result, 2, 64)
	if err != nil {
		fmt.Println("Failed to parse", result)
		fmt.Println(err)
	}
	return maskedVal
}

func mask2Int32(source int64, mask string) (addresses []int64) {
	result := ""
	input := strconv.FormatInt(int64(source), 2)
	for len(input) < 36 {
		input = "0" + input
	}
	for i, c := range mask {
		switch string(c) {
		case "0":
			result += string(input[i])
		case "1":
			result += "1"
		default:
			result += "X"
		}
	}
	allAddr := findAllAddr(result)
	for _, addr := range allAddr {
		intAddr, err := strconv.ParseInt(addr, 2, 64)
		if err != nil {
			fmt.Println("Failed to parse", addr)
			fmt.Println(err)
		}
		addresses = append(addresses, intAddr)
	}
	return addresses
}

func findAllAddr(addr string) (addresses []string) {
	for i, c := range addr {
		if string(c) == "X" {
			result1 := addr[:i] + "0" + addr[i+1:]
			result2 := addr[:i] + "1" + addr[i+1:]
			ans1 := findAllAddr(result1)
			ans2 := findAllAddr(result2)
			addresses = append(addresses, ans1...)
			addresses = append(addresses, ans2...)
			return addresses
		}
	}
	addresses = append(addresses, addr)
	return addresses
}
