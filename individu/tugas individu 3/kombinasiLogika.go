package main

import "fmt"

func main() {
	var p, q int
	fmt.Scan(&p, &q)

	fmt.Print((p % 2 == 0) || (q % 2 == 0), (p % 2 != 0) && (q % 2 != 0), !(p == q))
}