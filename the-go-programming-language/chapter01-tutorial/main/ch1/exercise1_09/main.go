package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s -- [%s]\n", b, resp.Status)
	}
}
