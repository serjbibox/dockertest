package apis

import (
	"fmt"
	"net/http"
)

var cnt int

func Hello(w http.ResponseWriter, r *http.Request) {
	cnt++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, guy! You are %d", cnt)
}
