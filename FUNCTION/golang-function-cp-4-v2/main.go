package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result []string
	for _, item := range data {
		if strings.Contains(item, input) {
			result = append(result, item)
		}
	}
	return strings.Join(result, ",")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
