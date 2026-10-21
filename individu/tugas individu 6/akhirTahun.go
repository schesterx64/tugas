package main
import "fmt"

func main() {
	var kartu, diskon, cashback, dapatKartu bool
	var totalBelanja, totalDiskon float64
	dapatKartu = false
	diskon = false
	cashback = false
	fmt.Scan(&totalBelanja)
	fmt.Scan(&kartu)
	
	if totalBelanja >= 100000 {
		diskon = true
	}
	if kartu == true && diskon == true {
		dapatKartu = true
	}
	if totalBelanja >= 200000 && dapatKartu == true {
		cashback = true
	}

	if diskon == true {
		totalDiskon = totalBelanja * 0.10
		totalBelanja -= totalDiskon
	}
	if cashback == true {
		totalBelanja -= 75000
	}


	fmt.Println("Kartu?", dapatKartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Printf("Rp. %.0f\n", totalBelanja)
}