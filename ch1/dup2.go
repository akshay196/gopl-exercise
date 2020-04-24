// dup2 Read input lines or text from number of files provided and
// display duplicate line inputted more than one with its count
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Enter text (Press Ctrl-d when done):")
		countLines(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error dup: %v\n", err)
				continue
			}
			countLines(f, count)
			f.Close()
		}
	}	
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	// No handled error occurred while reading file
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}	
}
