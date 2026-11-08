package main
import "fmt"

func main() {
	var input string
	A, B := 0, 0
	finishA, finishB := 0, 0
	pemenang := ""

	for i := 1; i <= 10; i++ {
		fmt.Scan(&input)

		if input == "A" {
			A++
			finishA++
			if finishA == 4 && pemenang == "" {
				pemenang = "A"
			}
		} else if input == "B" {
			B++
			finishB++
			if finishB == 4 && pemenang == "" {
				pemenang = "B"
			}
		}
	}
	if A != 5 || B != 5 {
		fmt.Println("tidak valid")
	} else {
		fmt.Println(pemenang)
	}
}