package main
import "fmt"

func main() {
	var poin int

	for {
		fmt.Scan(&poin)
		if poin >= 50 {
			break
		}
	}
	
	switch {
	case poin >= 200:
		fmt.Println("Gold user")
	case poin >= 100 && poin < 200:
		fmt.Println("Platinum user")
	case poin >= 50 && poin <= 99:
		fmt.Println("Silver user")
	default:
		fmt.Println("Poin tidak valid")
	}
}