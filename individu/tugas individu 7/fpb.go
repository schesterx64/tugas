package main
import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	
	for m != 0 {
		n, m = m, n % m
	}
	fmt.Println(n)
}