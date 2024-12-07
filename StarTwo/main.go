package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("locations.txt")
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanWords)
	var listA []int
	var listB []int
	flip := true
	for fileScanner.Scan() {
		locationStr := fileScanner.Text()
		location, err := strconv.ParseInt(locationStr, 10, 32)

		if err != nil {
			log.Printf("Error parsing location id: %s", err)
			return
		}

		if flip {
			flip = !flip
			listA = append(listA, int(location))
			continue
		}
		flip = !flip
		listB = append(listB, int(location))
	}

	simularityMap := make(map[int]int)

	for _, element := range listB {
		simularityMap[element] += 1
	}

	simularityScore := 0
	for _, element := range listA {
		simularityScore += element * simularityMap[element]
	}

	fmt.Printf("The simularity score is: %s\n", strconv.Itoa(simularityScore))

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
