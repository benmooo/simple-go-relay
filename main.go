package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the target URL
	targetURL, err := url.Parse("http:://localhost:44391")
	if err != nil {
		log.Fatal("Error parsing tart URL", err)
	}

	// Create a reverse proxy
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Start the proxy server listening on 0.0.0.0:3000
	http.Handle("/", reverseProxy)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
