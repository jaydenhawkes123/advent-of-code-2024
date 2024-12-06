package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open("locations.txt")
	if err != nil {
		log.Printf("Error reading file: %s", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		fmt.Printf("Location %s \n", fileScanner.Text())
	}
}
