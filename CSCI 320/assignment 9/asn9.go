package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

// Function to fetch the URL contents and print its length
func fetchURL(u string, wg *sync.WaitGroup) {
	defer wg.Done()

	// GET request to the URL
	resp, err := http.Get(u)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %s\n", u, err)
	} else {
		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading body: %s %s\n", u, err)
		} else {
			// Print the length of the response body
			fmt.Printf("Name and length: %s %d\n", u, len(body))
		}
	}
}

func main() {
	// Check if the required filename argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./asn9 <filename>")
		return
	}

	// Open the file containing URLs
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Error opening %s: %s\n", fn, err)
		return
	}
	defer f.Close()

	// Read the URLs from the file and store them in a slice
	urls := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	// Use a wait group to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Launch a goroutine for each URL
	for _, u := range urls {
		go fetchURL(u, &wg)
	}

	// Wait for all the goroutines to finish
	wg.Wait()
	// sleep for 5 seconds
	time.Sleep(5 * time.Second)
}
