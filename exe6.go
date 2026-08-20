package main

import (
	"fmt"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	APIkey := "secret123"
	apiKey := r.Header.Get("X-API-Key")

	if apiKey != APIkey {
		http.Error(w, "Authorization unsuccessful", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Welcome, authorization successful")
}
