package main

import (
	"bufio"
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
		location, err := strconv.Atoi(locationStr)

		if err != nil {
			log.Printf("Error parsing location id: %s", err)
			return
		}

		if flip {
			flip = !flip
			listA = append(listA, location)
			continue
		}
		flip = !flip
		listB = append(listB, location)
	}
}
