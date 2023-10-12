package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Initialize a map to count the frequency of each number
	frequency := make(map[int]int)

	// Read and analyze the data
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the CSV file:", err)
		return
	}

	for _, record := range records {
		for _, value := range record {
			num, err := strconv.Atoi(value)
			if err == nil {
				frequency[num]++
			}
		}
	}

	// Display the frequency of each number
	for number, count := range frequency {
		fmt.Printf("Number %d appeared %d times\n", number, count)
	}
}
