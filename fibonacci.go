package fibonacci

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// Response represents the JSON response structure
type Response struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

// fib calculates the n-th Fibonacci number using big.Int for large numbers.
func Fib(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

// FibHandler handles HTTP requests to the /fib endpoint.
func FibHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Get the 'n' query parameter
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Missing 'n' query parameter", http.StatusBadRequest)
		return
	}

	// Convert the query parameter to an integer
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		http.Error(w, "'n' must be a non-negative integer", http.StatusBadRequest)
		return
	}

	// Calculate the n-th Fibonacci number
	result := Fib(n)

	// Create response struct
	response := Response{
		Result: result.String(),
	}

	// Encode and write the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Handled request for n=%d in %v\n", n, time.Since(start))
}
