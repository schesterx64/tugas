package main
import "fmt"

func kuadran(x, y int) string {
	switch {
	case x == 0 && y == 0:
		return "Titik asal"
	case x == 0:
		return "Titik pada sumbu Y"
	case y == 0:
		return "Titik pada sumbu X"
	case x > 0 && y > 0:
		return "Kuadran I"
	case x < 0 && y > 0:
		return "Kuadran II"
	case x < 0 && y < 0:
		return "Kuadran III"
	case x > 0 && y < 0:
		return "Kuadran IV"
	default:
		return "Tidak valid"
	}
}

func main() {
	var x, y int
	fmt.Print("Koordinat X: ")
	fmt.Scan(&x)
	fmt.Print("Koordinat Y: ")
	fmt.Scan(&y)

	hasil := kuadran(x, y)
	fmt.Println("Hasil: ", hasil)
}