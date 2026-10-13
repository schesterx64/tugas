package main
import "fmt"

func main() {
	var x int
	var bilanganCantik bool
	fmt.Scan(&x)

	for x >= 10 {
		digitKanan := x % 10
		x /= 10
		digitKiri := x % 10

		if digitKiri > digitKanan {
			bilanganCantik = false
			break
		} else {
			bilanganCantik = true
		}
	}
	fmt.Println(bilanganCantik)
}