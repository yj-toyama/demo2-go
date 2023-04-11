package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hellogo", helloHandler)
	http.ListenAndServe(":8001", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// CORSヘッダーの追加
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "Hello, World!-GO")
}
