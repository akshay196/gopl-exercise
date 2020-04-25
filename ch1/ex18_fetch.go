// Fetch prints content from url
package main

import (
	"fmt"
	"io"
	"strings"
	"net/http"
	"os"
)

func main () {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") != true {
			url = "http://" + url
		}
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
