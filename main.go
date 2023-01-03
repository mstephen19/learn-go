package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"foo":"bar"}`))
		w.Header().Set("content-type", "application/json")
	}))
	http.ListenAndServe(":8000", nil)
}
