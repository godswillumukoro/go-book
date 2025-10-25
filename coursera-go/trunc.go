package main

import "fmt"

func main() {
	var pipe float32
	fmt.Println("What number is on your mind?")
	_, err := fmt.Scan(&pipe)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(int(pipe))
}