// Dup3 prints the count and text of lines that appear more than once
// across one or more input files, and shows the file names where each
// duplicated line appears.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filesPerLine := make(map[string]map[string]bool)

	// If no files are provided, exit gracefully.
	if len(os.Args[1:]) == 0 {
		fmt.Fprintln(os.Stderr, "usage: go run dup3.go file1.txt file2.txt ...")
		os.Exit(1)
	}

	// Loop through each file provided as an argument.
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", filename, err)
			continue
		}

		// Split file content into lines and record occurrences.
		for _, line := range strings.Split(string(data), "\n") {
			if line == "" {
				continue
			}
			counts[line]++
			addLine(filesPerLine, line, filename)
		}
	}

	// Print results: show only lines that appeared more than once.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("\tFound in: %s\n", strings.Join(getFileNames(filesPerLine[line]), ", "))
		}
	}
}

// addLine ensures each line has a map of filenames, then records the current file.
func addLine(filesPerLine map[string]map[string]bool, line, filename string) {
	if _, ok := filesPerLine[line]; !ok {
		filesPerLine[line] = make(map[string]bool)
	}
	filesPerLine[line][filename] = true
}

// getFileNames converts the inner map[string]bool to a slice of file names.
func getFileNames(fileMap map[string]bool) []string {
	files := make([]string, 0, len(fileMap))
	for file := range fileMap {
		files = append(files, file)
	}
	return files
}
