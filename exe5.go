package main

import (
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Header)
}
