package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			result += " "
		} else {
			if len(result) > 0 && result[len(result)-1] != ' ' {
				result += "_"
			}
			result += string(str[i])
		}
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
