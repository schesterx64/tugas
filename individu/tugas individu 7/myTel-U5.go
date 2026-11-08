package main
import "fmt"

func main() {
	totalGold, totalPlatinum, totalSilver := 0, 0, 0
	gold, platinum, silver := 0, 0, 0
	totalPoin := 0

	for totalPoin < 500 {
		var poin int

		for {
			fmt.Scan(&poin)
			if poin >= 50 {
				break
			}
		}
		totalPoin += poin
		
		switch {
		case poin > 200:
			totalGold++
			gold += poin
		case poin >= 100 && poin < 200:
			totalPlatinum++
			platinum += poin
		case poin >= 50 && poin <= 99:
			totalSilver++
			silver += poin
		}
	}
	var rataRataGold, rataRataPlatinum, rataRataSilver float64

	if totalGold > 0 {
		rataRataGold = float64(gold) / float64(totalGold)
	}
	if totalPlatinum > 0 {
		rataRataPlatinum = float64(platinum) / float64(totalPlatinum)
	}
	if totalSilver > 0 {
		rataRataSilver = float64(silver) / float64(totalSilver)
	}
	fmt.Printf("%.2f %.2f %.2f\n", rataRataGold, rataRataPlatinum, rataRataSilver)
}