package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/agent", agentHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/legacy", legacyHandler)
	//http.HandleFunc("/v2", v2Handler)
	fmt.Println("server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
