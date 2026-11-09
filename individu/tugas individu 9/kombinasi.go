package main
import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func C(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n - r))
}

func main() {
	var x, y, kombinasi int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	fx := faktorial(x)
	fy := faktorial(y)

	if x >= y {
		kombinasi = C(x, y)
	} else {
		kombinasi = C(y, x)
	}

	fmt.Println(fx, fy, kombinasi)
}