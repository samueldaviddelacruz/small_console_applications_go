package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `<!DOCTYPE html>
	<html>
	  <head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8" />
			<title>{{.Title}}</title>
	  </head>
	  <body>
			{{.Body}}

	
	  </body>
	</html>`
)

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	markdownText := flag.String("text", "", "Markdown text to preview")
	skipPreview := flag.Bool("s", false, "Skip opening the preview in the browser")
	tFname := flag.String("t", "", "Alternate template file name")
	flag.Parse()
	// if user did not provide input file, show usage

	if *filename == "" && *markdownText == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *markdownText != "" {
		// if user provided markdown text, write to temp file
		temp, err := os.CreateTemp("", "mdp*.md")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if _, err := temp.WriteString(*markdownText); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := temp.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		*filename = temp.Name()
		defer os.Remove(*filename)
	}
	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, tFname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)
	if saveHTML(outName, htmlData); err != nil {
		return err
	}
	if skipPreview {
		return nil
	}
	defer os.Remove(outName)
	return preview(outName)
}

func saveHTML(outName string, data []byte) error {
	// Write the bytes to the output file
	return os.WriteFile(outName, data, 0644)
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML file
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Create a new template and parse the default template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if os.Getenv("DEFAULT_TEMPLATE") != "" {
		t, err = template.New("mdp").Parse(os.Getenv("DEFAULT_TEMPLATE"))
	}
	if err != nil {
		return nil, err
	}
	// if user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}
	// Create a new content struct and populate it with the
	// title and body of the markdown file
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	//Create a buffer of bytes to write to the output file
	var buffer bytes.Buffer

	// execute the template and write the output to the buffer
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}
	// Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd"
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("unsupported platform")
	}
	// Append filename to params slice
	cParams = append(cParams, fname)
	// Locate executable in path
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}
	// open the file using default program
	err = exec.Command(cPath, cParams...).Run()
	// Give the browser some time to open the file before deleting
	time.Sleep(2 * time.Second)
	return err
}
