package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/joho/godotenv"
)

func main() {

	// Load variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	remote, err := url.Parse("http://localhost:44391")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
