package fibonacci

import (
    "math/big"
    "testing"
)

// Teste para a função Fib
func TestFibonacci(t *testing.T) {
    testCases := []struct {
        n        int
        expected *big.Int
    }{
        {0, big.NewInt(0)},
        {1, big.NewInt(1)},
        {2, big.NewInt(1)},
        {3, big.NewInt(2)},
        {4, big.NewInt(3)},
        {5, big.NewInt(5)},
        {6, big.NewInt(8)},
        {20, big.NewInt(6765)}, // Caso de teste adicional
        {50, big.NewInt(12586269025)}, // Outro caso de teste adicional
    }

    for _, tc := range testCases {
        result := Fib(tc.n)
        if result.Cmp(tc.expected) != 0 {
            t.Errorf("For n=%d, expect %s, received %s", tc.n, tc.expected.String(), result.String())
        }
    }
}
