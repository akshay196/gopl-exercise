// dup1 Read input lines and display duplicate line inputted more than
// one with its count
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (Press Ctrl-d when done):")
	for input.Scan() {
		count[input.Text()]++
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
