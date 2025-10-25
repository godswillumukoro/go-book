// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numbers = make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter an integer, or X to quit.")

	for {
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()

		if text == "X" {
			break
		}

		n, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or X to quit.")
			continue
		}
		numbers = append(numbers, n)
		sort.Ints(numbers)
		fmt.Println("Sorted slice: ", numbers)
	}
	// sort slice
	fmt.Println("Final slice: ", numbers)
}