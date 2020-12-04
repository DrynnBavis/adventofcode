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
	input, err := ioutil.ReadFile("04_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse into slice of raw lines
	s := strings.Split(string(input), "\n")
	passports := []string{}
	newLine := true
	count := 0
	for _, line := range s {
		if line == "" {
			newLine = true
			count++
		} else {
			if newLine {
				passports = append(passports, line)
			} else {
				passports[count] += " " + line
			}
			newLine = false
		}
	}

	// Question #8
	numValid := question8(passports)
	fmt.Println(numValid)
}

func question8(input []string) (numValid int) {
	requiredFields := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	} //"cid" removed
	passportMaps := []map[string]string{}
	for _, passport := range input {
		entries := strings.Split(passport, " ")
		m := make(map[string]string)
		for _, entry := range entries {
			kv := strings.Split(entry, ":")
			key, val := string(kv[0]), string(kv[1])
			m[key] = val
		}
		passportMaps = append(passportMaps, m)
		correctCount := 0
		if val, ok := m["byr"]; ok {
			i, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				if i >= 1920 && i <= 2002 {
					correctCount++
				}
			}
		}
		if val, ok := m["iyr"]; ok {
			i, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				if i >= 2010 && i <= 2020 {
					correctCount++
				}
			}
		}
		if val, ok := m["eyr"]; ok {
			i, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				if i >= 2020 && i <= 2030 {
					correctCount++
				}
			}
		}
		if val, ok := m["hgt"]; ok {
			unit := val[len(val)-2:]
			i, err := strconv.ParseInt(val[:len(val)-2], 10, 64)
			if err == nil {
				if unit == "cm" {
					if i >= 150 && i <= 193 {
						correctCount++
					}
				}
				if unit == "in" {
					if i >= 59 && i <= 76 {
						correctCount++
					}
				}
			}
		}
		if val, ok := m["hcl"]; ok {
			if string(val[0]) != "#" || len(val) != 7 {
				continue
			}
			_, err := strconv.ParseInt(val[1:], 16, 64)
			if err == nil {
				correctCount++
			}
		}
		if val, ok := m["ecl"]; ok {
			validEyeColours := map[string]bool{
				"amb": true,
				"blu": true,
				"brn": true,
				"gry": true,
				"grn": true,
				"hzl": true,
				"oth": true,
			}
			if _, ok := validEyeColours[string(val)]; ok {
				correctCount++
			}
		}
		if val, ok := m["pid"]; ok {
			if len(string(val)) != 9 {
				continue
			}
			_, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				correctCount++
			}
		}
		if correctCount == len(requiredFields) {
			numValid++
		}
	}
	return numValid
}
