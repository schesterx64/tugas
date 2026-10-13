package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	total := 0
	

	for x > 0 {
		digit := x % 10
		kuadrat := digit * digit

		fmt.Printf(" %d", kuadrat)
	total += kuadrat
	x /= 10
	}
	fmt.Println()
	fmt.Println(total)
}