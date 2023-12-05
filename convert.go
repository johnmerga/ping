package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func filter(inputFileName string) ([]string, error) {
	ips := make([]string, 0)
	outputFileName := strings.TrimSuffix(inputFileName, ".csv") + "_filtered.csv"

	csvFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return nil, err
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil, err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)

	for _, record := range records {
		if len(record) > 0 {
			split := strings.Split(record[0], "\t")
			name := split[0]
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

	writer.Flush()

	// Check for errors during flushing
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return nil, err
	}
	fmt.Println("Extraction and saving complete.")
	return ips, nil
}
