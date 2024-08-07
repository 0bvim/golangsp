package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCsvFile(filePath string) ([][]string, error) {
	//open just give the reference of content in memory.
	file, err := os.Open(filePath)
	//to close file if when function ends with error or success.
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error opening file %w", err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records %w", err)
	}
	return records, nil
}
