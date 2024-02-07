package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"jsbeautifier-go/jsbeautifier"
	"net/http"
	"time"
)

func handleErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	
	URL := flag.String("u", "", "URL to fetch")
	flag.Parse()

	client := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	if *URL == "" {
		fmt.Println("Please provide a URL using the -u flag.")
		return
	}

	resp, err := client.Get(*URL)
	handleErr(err)
	

	defer func() {
		if resp != nil && resp.Body != nil {
			if err := resp.Body.Close(); err != nil {
				fmt.Printf("Error closing response body for => %s: %v\n", *URL, err)
			}
		}
	}()

	jsContent, err := io.ReadAll(resp.Body)

	handleErr(err)

	jsString := string(jsContent)

	opts := jsbeautifier.DefaultOptions()
	out, err := jsbeautifier.Beautify(&jsString, opts)
	handleErr(err)
	fmt.Println(out)


}
