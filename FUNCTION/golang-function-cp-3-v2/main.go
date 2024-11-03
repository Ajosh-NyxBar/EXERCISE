package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	separators := []string{" ", ",", ";"}
	for _, sep := range separators {
		names = strings.ReplaceAll(names, sep, " ")
	}
	nameList := strings.Fields(names)
	shortestName := nameList[0]
	for _, name := range nameList {
		if len(name) < len(shortestName) || (len(name) == len(shortestName) && name < shortestName) {
			shortestName = name
		}
	}

	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
