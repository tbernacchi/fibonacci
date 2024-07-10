package main

import (
	"fmt"
	"net/http" 
	"fibonacci"
)

func main() {
	// Set up the /fib endpoint
	http.HandleFunc("/fib", fibonacci.FibHandler)

	// Start the server on port 8000
	fmt.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
