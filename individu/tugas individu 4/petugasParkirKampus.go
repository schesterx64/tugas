package main

import "fmt"

func main() {
	var kapasitasMaksimum = 50
	var kendaraan int

	for kendaraan = 0; kendaraan <= kapasitasMaksimum; kendaraan++ {
		fmt.Printf("Jumlah kendaraan saat ini: %d\n", kendaraan)
	} 
	fmt.Println("Parkir Penuh!")
}