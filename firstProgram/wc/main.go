package main

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func main() {

	//defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "count lines")
	//defining a boolean flag -b to count bytes instead of words
	bytes := flag.Bool("b", false, "count bytes")
	//parsing the flags provided by the user
	flag.Parse()
	//calling the count function to count the number of words or lines
	println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text  from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (defaut is split by lines)
	// This is done only if the user did not provide the -l flag
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	if countLines {
		scanner.Split(bufio.ScanLines)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}
	// defining a counter
	wc := 0
	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}
