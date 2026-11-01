package main
import "fmt"

func main() {
	total := 0

	for i := 1; i <= 5; i++ {
		var n int
		for {
			fmt.Printf("Hari %d: ", i)
			fmt.Scan(&n)

			if n >= 0 && n <= 200 {
				break
			}
		}
		total += n
	}
	fmt.Printf("Jumlah pengunjung: %d orang \n", total)
}