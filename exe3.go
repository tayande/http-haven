package main

import (
	"fmt"
	"io"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		data, _ := io.ReadAll(r.Body)
		text := string(data)
		count := len(text)
		fmt.Fprintln(w, count)
	}
}
