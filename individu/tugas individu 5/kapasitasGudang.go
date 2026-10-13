package main
import "fmt"

func main() {
	var K, V int
	fmt.Scan(&K)
	totalBarang := 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&V)
		totalBarang += V

		//fmt.Println(totalBarang >= K)
	}
	fmt.Println(totalBarang >= K)
} 