// echo4
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit new line")
var s = flag.String("s", " ", "Separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
