package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func([]float64) float64

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2Float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV reader used to read in data from csv files
	cr := csv.NewReader(r)
	cr.ReuseRecord = true
	// Adjusting for 0 based index
	column--
	// Read all csv data

	var data []float64

	// looping through all records
	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Cannot read data from file: %w", err)
		}
		// Skip the header row
		if i == 0 {
			continue
		}
		// Check if the column is valid
		if len(row) <= column {
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}
		// try to Convert the data to a float64
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	// return the slice of float64 and nil error
	return data, nil
}
