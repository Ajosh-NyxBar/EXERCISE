package main

import (
	"fmt"
)

// CountingNumber menjumlahkan angka dari 1 hingga n dengan peningkatan 0.5
func CountingNumber(n int) float64 {
	var sum float64
	for i := 1.0; i <= float64(n); i += 0.5 {
		sum += i
	}
	return sum
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println(CountingNumber(10))  // Output: 104.5
	fmt.Println(CountingNumber(100)) // Output: 10049.5
}
