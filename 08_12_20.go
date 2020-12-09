package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	cmd  string
	sign string
	val  int
}

func main() {
	// Read input
	input, err := ioutil.ReadFile("08_12_20.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions := []Instruction{}
	s := strings.Split(string(input), "\n")
	for _, line := range s {
		k := strings.Split(string(line), " ")
		cmd := k[0]
		sign := string(k[1][0])
		val, err := strconv.ParseInt(string(k[1][1:]), 10, 64)
		if err != nil {
			fmt.Println("Failed to parse ", string(string(k[1][1:])))
		}
		newInstruction := Instruction{cmd, sign, int(val)}
		instructions = append(instructions, newInstruction)
	}

	// Question #1:
	fmt.Println(findAccBeforeLoop(instructions))

	// Question #2:
	fmt.Println(findAccAfterPatch(instructions))
}

func findAccBeforeLoop(instructions []Instruction) (acc int) {
	seen := map[int]bool{}
	i := 0
	for {
		// have we been here before?
		if _, ok := seen[i]; ok {
			// break loop if we've been on this index before
			break
		} else {
			// otherwise add to seen
			seen[i] = true
		}

		// process this instruction
		instruction := instructions[i]
		switch instruction.cmd {
		case "jmp":
			if instruction.sign == "+" {
				i += instruction.val
			} else {
				i -= instruction.val
			}
		case "acc":
			if instruction.sign == "+" {
				acc += instruction.val
			} else {
				acc -= instruction.val
			}
			i++
		default:
			// must be nop
			i++
		}
	}
	return acc
}

func findAccAfterPatch(instructions []Instruction) (acc int) {
	seen := map[int]bool{}
	spotsToCheck := []int{}
	history := []int{}
	i := 0
	backtracking := false
	shouldSwap := false
	best := 0
	for {
		// have we been here before?
		if _, ok := seen[i]; ok {
			// we have, let's backtrack
			backtracking = true
			history = history[:len(history)-1]
			i = history[len(history)-1]
			for i != spotsToCheck[len(spotsToCheck)-1] {
				instruction := instructions[i]
				delete(seen, i)
				history = history[:len(history)-1]
				if instruction.cmd == "acc" {
					if instruction.sign == "+" {
						acc -= instruction.val
					} else {
						acc += instruction.val
					}
				}
				i = history[len(history)-1]
			}
			spotsToCheck = spotsToCheck[:len(spotsToCheck)-1]
			shouldSwap = true
		} else {
			// otherwise add to seen
			seen[i] = false
		}

		// process this instruction
		instruction := instructions[i]
		switch instruction.cmd {
		case "jmp":
			if !backtracking {
				spotsToCheck = append(spotsToCheck, i)
			}
			if shouldSwap {
				shouldSwap = false
				i++
			} else {
				if instruction.sign == "+" {
					i += instruction.val
				} else {
					i -= instruction.val
				}
			}
		case "acc":
			if instruction.sign == "+" {
				acc += instruction.val
			} else {
				acc -= instruction.val
			}
			i++
		default:
			// must be nop
			if !backtracking {
				spotsToCheck = append(spotsToCheck, i)
			}
			if shouldSwap {
				shouldSwap = false
				if instruction.sign == "+" {
					i += instruction.val
				} else {
					i -= instruction.val
				}
			} else {
				i++
			}
		}

		history = append(history, i)
		if i > best {
			best = i
		}

		// are we at EoF?
		if i == len(instructions) {
			break
		}
	}
	return acc
}
