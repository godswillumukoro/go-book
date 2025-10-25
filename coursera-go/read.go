// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Name struct {
	fname string
	lname string
}

func main() {
	var names []Name


	fmt.Println("Name of the text file please: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if line == "" {
			continue
		}

		token := strings.Fields(line)
		if len(token) ==2 {
			n := Name{
				fname: token[0],
				lname: token[1],
			}
			names = append(names, n)
		} else {
			fmt.Println("Skipping invalid line: ", line)
		}
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("\nNames read from file:")
	for _, n := range names {
		fmt.Printf("First: %-10s Last: %-10s\n", n.fname, n.lname)
	}

}