package main
import "fmt"

func main() {
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
		case poin >= 200:
			gold++
		case poin >= 100 && poin < 200:
			platinum++
		case poin >= 50 && poin <= 99:
			silver++
		}
	}
	fmt.Println("Gold user:", gold)
	fmt.Println("Platinum user:", platinum)
	fmt.Println("Silver user:", silver)
}