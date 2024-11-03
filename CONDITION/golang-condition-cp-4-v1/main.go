package main

import "fmt"

func GetTicketPrice(VIP int, regular int, student int, day int) float32 {
	const vipPrice = 30
	const regularPrice = 20
	const studentPrice = 10

	
	totalPrice := float32(VIP*vipPrice + regular*regularPrice + student*studentPrice)


	if totalPrice < 100 {
		return totalPrice
	}


	totalTickets := VIP + regular + student


	var discount float32
	if day%2 == 1 { 
		if totalTickets < 5 {
			discount = 0.15
		} else {
			discount = 0.25
		}
	} else { // Tanggal genap
		if totalTickets < 5 {
			discount = 0.10
		} else {
			discount = 0.20
		}
	}


	discountedPrice := totalPrice * (1 - discount)
	return discountedPrice
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
