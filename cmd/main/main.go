package main

import (
	"net/http"

	H "neighborhost_service/api"
)

func main() {
	http.HandleFunc("/api", H.Handler)
	http.ListenAndServe(":8080", nil)
}
