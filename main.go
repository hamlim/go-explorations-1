package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Success bool `json:"success"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-foo", "bar")
		fmt.Fprintf(w, "<h1>Hello, World!</h1>")
		// w.Write()
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := Data{Success: true}
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}
