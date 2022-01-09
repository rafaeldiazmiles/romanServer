package main

import (
	"net/http"
	"strings"
)

func main() {
	// htttp package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {

		}
	})
}
