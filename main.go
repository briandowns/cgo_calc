package main

import (
	"fmt"

	"github.com/briandowns/cgo_calc/calc"
)

// Think of this as the Kubernetes code calling functions in the Go
// standard library.

func main() {
	fmt.Printf("Addition:       10 + 10     = %.2f\n", calc.Add(10, 10))
	fmt.Printf("Subtraction:    98 - 66   = %.2f\n", calc.Sub(98, 66))
	fmt.Printf("Multiplication: 100 * -18 = %.2f\n", calc.Mul(100, -18))
	fmt.Printf("Division:       12 / 4    = %.2f\n", calc.Div(12, 4))
}
