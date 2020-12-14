package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("13_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")
	earliest, err := strconv.ParseInt(string(s[0]), 10, 64)
	if err != nil {
		fmt.Println("Failed to parse ", s[0])
	}
	busses := strings.Split(s[1], ",")

	// Question #1:
	nextBus, nextTime := findEarliestBusAndTime(int(earliest), busses)
	fmt.Println(nextTime * nextBus)

	// Question #2:
	fmt.Println(findTimeForSequence(busses))
}

func findEarliestBusAndTime(earliest int, busses []string) (busID int, bestTime int) {
	bestTime = earliest
	for _, bus := range busses {
		if bus == "x" {
			continue
		}
		busTime, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Println("Failed to parse ", bus)
		}
		timeTilNext := busTime - (earliest % busTime)
		if timeTilNext < bestTime {
			busID = busTime
			bestTime = timeTilNext
		}
	}
	return busID, bestTime
}

func findTimeForSequence(busses []string) (time int) {
	indexes := []int{}
	primes := []int{}
	for i, bus := range busses {
		if bus == "x" {
			continue
		}
		busTime, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Println("Failed to parse ", bus)
		}
		indexes = append(indexes, i)
		primes = append(primes, busTime)
	}
	x := 0
	for i, prime := range primes {
		if prime > primes[x] {
			x = i
		}
	}
	productPrimes := 1
	for _, prime := range primes {
		productPrimes *= prime
	}
	for k := 1; k < productPrimes; k++ {
		time = k*primes[x] - indexes[x]
		foundSoln := true
		for i := 0; i < len(indexes); i++ {
			if (time+indexes[i])%primes[i] != 0 {
				foundSoln = false
				break
			}
		}
		if foundSoln {
			break
		}
	}

	return time
}

func solveSysOfCongruent(remainders []int, coprimes []int) (result int) {
	fmt.Println(remainders)
	fmt.Println(coprimes)
	m := []int{}
	productCoprimes := 1
	for _, coprime := range coprimes {
		productCoprimes *= coprime
	}
	for x := range remainders {
		m = append(m, 1)
		for y, coprime := range coprimes {
			if x == y {
				continue
			}
			m[x] = m[x] * coprime
		}
	}
	y := []int{}
	for x := range remainders {
		y = append(y, 0)
		for i := 1; i <= coprimes[x]; i++ {
			if (i*m[x])%coprimes[x] == remainders[x]-1 {
				y[x] = i
				break
			}
		}
	}
	for x := range remainders {
		r2 := coprimes[x] - remainders[x]
		fmt.Println(r2, m[x], y[x])
		result += r2 * m[x] * y[x]
	}
	return result % productCoprimes
}
