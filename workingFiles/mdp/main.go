package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
	  <head>
	    <meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>Markdown Preview Tool</title>
	</head>
	 <body>`
	footer = `</body>
	</html>`
)

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()
	// if user did not provide input file, show usage

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, out io.Writer) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	htmlData := parseContent(input)
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)
	return saveHTML(outName, htmlData)
}

func saveHTML(outName string, data []byte) error {
	// Write the bytes to the output file
	return os.WriteFile(outName, data, 0644)
}

func parseContent(input []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML file
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	//Create a buffer of bytes to write to the output file
	var buffer bytes.Buffer

	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}
