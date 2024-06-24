package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api", Handler)
	http.ListenAndServe(":8080", nil)
}
