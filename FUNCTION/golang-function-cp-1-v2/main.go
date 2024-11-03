package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	dayStr := fmt.Sprintf("%02d", day)
	monthStr := months[month-1]

	return fmt.Sprintf("%s-%s-%d", dayStr, monthStr, year)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
