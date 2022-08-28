package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./hello-world <arguments>\n")
		os.Exit(1)
	}
	// fmt.Printf("hello world\n1st Argument: %v\n", args[1])
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s", response.StatusCode, body)
}
