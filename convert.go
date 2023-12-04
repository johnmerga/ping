package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func filter(inputFileName string) ([]string, error) {
	ips := make([]string, 0)
	// Replace "input.csv" and "output.txt" with your actual file names
	outputFileName := "output.txt"

	// Open the CSV file
	csvFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil, err
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return nil, err
	}

	// Open the output file for writing
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil, err
	}
	defer outputFile.Close()

	// Create a writer for the output file
	writer := csv.NewWriter(outputFile)

	// Iterate over each record and write the first column to the output file
	for _, record := range records {
		if len(record) > 0 {
			// Write the first column to the output file
			// split by tab
			split := strings.Split(record[0], "\t")
			// take the first element
			name := split[0]
			// check if the string contain  ".siinqeebank.com"
			if !strings.Contains(name, ".siinqeebank.com") {
				append := ".siinqeebank.com"
				name = name + append
			}
			ips = append(ips, name)
			err := writer.Write([]string{name})
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return nil, err
			}
		}
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	// Check for errors during flushing
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return nil, err
	}
	fmt.Println("Extraction and saving complete.")
	return ips, nil
}
