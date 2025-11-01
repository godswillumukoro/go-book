// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

// Submit your Go program source code.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write a list of 10 random numbers delimited by comma. Eg. 1,2,3")
	scanner.Scan()
	text := scanner.Text()
	tempNum := strings.Split(text, ",")
	fmt.Println("tempNum", tempNum)
	for i:=0; i<len(tempNum); i++ {
		// println(i)
			convNum, err := strconv.Atoi(tempNum[i])
			if err != nil {
				fmt.Println("Incorrect value passed.")
			}
			num = append(num, convNum)
	}
	BubbleSort(num)
	fmt.Println("Sorted: ", num)
}

func BubbleSort(s []int) {
	for i:=0; i<len(s); i++{
		for j:=0; j<len(s)-1; j++{
			if s[j]>s[j+1] {
				Swap(s,j)
			}
		}
	}
}

func Swap(s []int, j int) {
	s[j], s[j+1] = s[j+1], s[j]
}