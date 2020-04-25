// Fetch prints content from url
package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
)

func main () {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err1 := io.Copy(os.Stdout, resp.Body)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: Reading %v\n", err)
			os.Exit(1)
		}
	}
}
