package main
import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	pertama := true
	limit := n

	if m < n {
		limit = m
	}

	for i := 1; i <= limit; i++ {
		if n % i == 0 && m % i == 0 {
			if !pertama {
				fmt.Print(", ")
			}
			fmt.Print(i)
			pertama = false
		}
	}
	fmt.Println()
}