package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong get method", http.StatusMethodNotAllowed)
		return
	}
	//var result int
	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	numA, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	numB, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "request not accepted", http.StatusBadRequest)
		return
	}
	var result int
	if op == "add" {
		result = numA + numB
	}
	if op == "subtract" {
		result = numA - numB
	}
	if op == "multiply" {
		result = numA * numB
	}
	if op == "divide" && numB == 0 {
		http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		return
	}
	if op == "divide" {
		result = numA / numB
	}

	if op != "add" && op != "subtract" && op != "multiply" && op != "divide" {
		http.Error(w, "wrong choices", http.StatusBadRequest)
		return
	}
	fmt.Println("op:", op)
	fmt.Println("numA:", numA)
	fmt.Println("numB:", numB)
	fmt.Println("result:", result)

	fmt.Fprintln(w, "result:", result)

}

 