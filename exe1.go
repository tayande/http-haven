package main

import (
	"fmt"
	"net/http"
)


func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong get method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "pong")
}