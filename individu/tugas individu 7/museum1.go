package main
import "fmt"

func main() {
	var n int
	total := 0
	jumlah := 0
	kenaikan := 0
	s := -1
	
	for {
		fmt.Scan(&n)

		if n < 0 || n > 200 {
			break
		}
		if s != -1 && n > s {
			kenaikan++
		}
		total += n
		jumlah++
		s = n
	}
	var rataRata float64
	if jumlah > 0 {
		rataRata = float64(total) / float64(jumlah)
	} else {
		rataRata = 0
	}
	if rataRata == float64(int(rataRata)) {
		fmt.Printf("%d %.0f\n", kenaikan, rataRata)
	} else {
		fmt.Printf("%d %.2f\n", kenaikan, rataRata)
	}
}