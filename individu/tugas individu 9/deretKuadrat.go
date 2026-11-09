package main
import "fmt"

func deretKuadrat(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	hasil := deretKuadrat(n)
	fmt.Println(hasil)
}