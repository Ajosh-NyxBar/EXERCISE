package main

import "fmt"

// BMICalculator menghitung BMI berdasarkan gender dan tinggi badan
func BMICalculator(gender string, height int) float64 {
	base := float64(height - 100)
	if gender == "laki-laki" {
		return base - (base * 0.10)
	} else if gender == "perempuan" {
		return base - (base * 0.15)
	}
	return 0.0
}

func main() {
	fmt.Println(BMICalculator("laki-laki", 170))  // Output: 63
	fmt.Println(BMICalculator("perempuan", 165)) // Output: 55.25
}
