package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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

	slices.SortFunc(listA, func(i, j int) int {
		return cmp.Compare(i, j)
	})
	slices.SortFunc(listB, func(i, j int) int {
		return cmp.Compare(i, j)
	})
	totalDistance := 0
	for index := range listA {
		totalDistance += absInt(listA[index] - listB[index])
	}

	fmt.Printf("The total difference is: %s\n", strconv.Itoa(totalDistance))
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
