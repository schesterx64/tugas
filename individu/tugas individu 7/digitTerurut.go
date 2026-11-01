package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	digitTerakhir := n % 10
	n = n / 10
	ascending := true
	descending := true

	for n > 0 {
		digitS := n % 10
		if digitS > digitTerakhir {
			ascending = false
		} else if digitS < digitTerakhir {
			descending = false
		}
		digitTerakhir = digitS
		n = n / 10
	}

	if ascending == true {
		fmt.Println("Ascending")
	} else if descending == true {
		fmt.Println("Descending")
	} else {
		fmt.Println("Tidak terurut")
	}
}