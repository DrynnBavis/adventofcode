package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read Rules
	rulesInput, err := ioutil.ReadFile("16_12_20_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	s1 := strings.Split(string(rulesInput), "\n")
	rules := map[int][]string{}
	for _, line := range s1 {
		split1 := strings.Split(line, ": ")
		rule := split1[0]
		split2 := strings.Split(split1[1], " or ")
		split3a := strings.Split(split2[0], "-")
		split3b := strings.Split(split2[1], "-")
		range1Start, _ := strconv.Atoi(split3a[0])
		range1End, _ := strconv.Atoi(split3a[1])
		range2Start, _ := strconv.Atoi(split3b[0])
		range2End, _ := strconv.Atoi(split3b[1])
		for i := range1Start; i <= range1End; i++ {
			if _, ok := rules[i]; ok {
				rules[i] = append(rules[i], rule)
			} else {
				rules[i] = []string{rule}
			}
		}
		for i := range2Start; i <= range2End; i++ {
			if _, ok := rules[i]; ok {
				rules[i] = append(rules[i], rule)
			} else {
				rules[i] = []string{rule}
			}
		}
	}

	// Read other Tickets
	otherTicketsInput, err := ioutil.ReadFile("16_12_20_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	s3 := strings.Split(string(otherTicketsInput), "\n")
	otherTicketVals := []int{}
	ticketScanningErrorRate := 0
	for _, line := range s3 {
		vals := strings.Split(line, ",")
		ticketVals := []int{}
		for _, val := range vals {
			valInt, _ := strconv.Atoi(val)
			if _, ok := rules[valInt]; !ok {
				ticketScanningErrorRate += valInt
				break
			} else {
				ticketVals = append(ticketVals, valInt)
			}
		}
		if len(ticketVals) == 20 {
			otherTicketVals = append(otherTicketVals, ticketVals...)
		}
	}
	fmt.Println(ticketScanningErrorRate)

	// Read My Ticket
	myTicketInput, err := ioutil.ReadFile("16_12_20_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	myTicket := strings.Split(string(myTicketInput), ",")
	numOtherTickets := len(otherTicketVals) / len(myTicket)
	numFields := len(myTicket)

	// Find all possible fields per column
	columnToFields := map[int][]string{}
	for i := 0; i < 20; i++ {
		m := map[string]int{}
		for k := 0 + i; k < len(otherTicketVals); k += 20 {
			myRules := rules[otherTicketVals[k]]
			for _, rule := range myRules {
				if _, ok := m[rule]; !ok {
					m[rule] = 1
				} else {
					m[rule]++
				}
			}
		}
		for field, occurences := range m {
			if occurences == numOtherTickets {
				// if _, ok := fieldColumns[field]; !ok {
				// 	columnToFields[i] = append(columnToFields[i], field)
				// }
				columnToFields[i] = append(columnToFields[i], field)
			}
		}
	}
	discoveredFields := []string{}
	ColToField := map[int]string{}
	for {
		if len(discoveredFields) == numFields {
			break
		}
		for col, rules := range columnToFields {
			remainder := difference(rules, discoveredFields)
			if len(remainder) == 1 {
				discoveredFields = append(discoveredFields, remainder[0])
				columnToFields[col] = remainder
				ColToField[col] = remainder[0]
			}
		}
	}
	finalAnswer := 1
	for k, v := range ColToField {
		splits := strings.Split(string(v), " ")
		if splits[0] == "departure" {
			val, _ := strconv.Atoi(myTicket[k])
			finalAnswer *= val
		}
	}
	fmt.Println(finalAnswer)

}

func difference(main []string, removal []string) []string {
	remainder := []string{}

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for _, s1 := range main {
		found := false
		for _, s2 := range removal {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			remainder = append(remainder, s1)
		}
	}

	return remainder
}
