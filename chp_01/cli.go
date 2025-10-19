// Prints it command line arguments
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func explicitLoop() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func implicitLoop() {
	s, sep := "", ""

	for _, arg := range(os.Args[1:]) {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func stringJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func simplePrint() {
	fmt.Println(os.Args[1:])
}

func measureTime(tag string, f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %v\n", tag, elapsed)
}

func main() {
	measureTime("Explicit loop", explicitLoop)
	measureTime("Implicit loop", implicitLoop)
	measureTime("strings.Join", stringJoin)
	measureTime("Simple print", simplePrint)
}