package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	roman "github.com/rafaeldiazmiles/romanServer/romanNumerals"
)

func main() {
	// htttp package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is no in the list, send Not Found Status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found - hoolis\n"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(roman.Numerals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request- maloo\n"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
