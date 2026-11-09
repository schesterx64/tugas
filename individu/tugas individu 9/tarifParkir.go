package main
import "fmt"

func tarif(jam int) int {
	if jam <= 2 {
		return 5000
	}
	return 5000 + (jam - 2) * 3000
}

func main() {
	var lamaParkir int
	fmt.Print("Lama parkir: ")
	fmt.Scan(&lamaParkir)

	total := tarif(lamaParkir)
	fmt.Println("Total tarif parkir: Rp.", total)
}