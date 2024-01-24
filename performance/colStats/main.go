package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

func main() {
	// Verify and parse arguments
	op := flag.String("op", "sum", "Operation to perform on the data")
	column := flag.Int("col", 1, "CSV column on which to execute the operation")
	flag.Parse()
	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, column int, out io.Writer) error {
	var opFunc statsFunc
	if len(filenames) == 0 {
		return ErrNoFiles
	}
	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}
	// Validate the operation and define the opFunc accordingly
	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}
	consolidate := make([]float64, 0)

	// Create the channel to receive the results or errors of operations
	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	filesCh := make(chan string)

	wg := sync.WaitGroup{}

	// Loop through all files sending them to the files channel
	// so each one will be processed when a worker is available
	go func() {
		defer close(filesCh)
		for _, filename := range filenames {
			filesCh <- filename
		}
	}()

	// Loop through all the files and create a goroutine for each
	for i := 0; i < runtime.NumCPU(); i++ {
		// Increment the WaitGroup counter
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fname := range filesCh {
				// Open the file for reading
				f, err := os.Open(fname)
				if err != nil {
					errCh <- fmt.Errorf("cannot open file: %w", err)
					return
				}
				// Parse the CSV into a slice of float64 numbers
				data, err := csv2Float(f, column)
				if err != nil {
					errCh <- err
				}
				if err := f.Close(); err != nil {
					errCh <- err
				}
				// Send the data to the result channel
				resCh <- data
			}
		}()
	}
	go func() {
		// Wait for all the goroutines to finish
		wg.Wait()
		close(doneCh)
	}()
	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidate))
			return err
		}
	}
}
