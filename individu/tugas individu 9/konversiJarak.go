package main
import "fmt"

func konversi(km, m, cm int) int {
	return (km * 100000) + (m * 100) + cm
}

func main() {
	var km, m, cm int
	fmt.Print("Jarak dalam km: ")
	fmt.Scan(&km)
	fmt.Print("Jarak dalam m: ")
	fmt.Scan(&m)
	fmt.Print("Jarak dalam cm: ")
	fmt.Scan(&cm)

	total := konversi(km, m, cm)
	fmt.Printf("Hasil konversi: %d cm", total)
}