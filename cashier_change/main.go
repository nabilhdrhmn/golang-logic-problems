package main

import (
	"fmt"
)

func main() {
	var totalBelanja, uangDibayar int
	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Pembeli membayar: Rp ")
	fmt.Scan(&uangDibayar)

	change := uangDibayar - totalBelanja
	if change < 0 {
		fmt.Println("False, kurang bayar")
		return
	}

	// Display unrounded change
	fmt.Printf("Kembalian yang harus diberikan kasir: %d,\n", change)

	// Round down to the nearest 100
	roundedChange := (change / 100) * 100
	fmt.Printf("dibulatkan menjadi %d\n", roundedChange)

	// Show denominations
	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	for _, p := range pecahan {
		if roundedChange >= p {
			lembar := roundedChange / p
			fmt.Printf("%d lembar %d\n", lembar, p)
			roundedChange %= p
		}
	}
}
