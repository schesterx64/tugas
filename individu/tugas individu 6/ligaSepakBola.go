package main
import "fmt"

func main() {
	var gol1, gol2, gol3, gol4 int
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)
	golBesar := gol1
	golKecil := gol1

	if gol2 > golBesar {
		golBesar = gol2
	}
	if gol3 > golBesar {
		golBesar = gol3
	}
	if gol4 > golBesar {
		golBesar = gol4
	}

	if gol2 < golKecil {
		golKecil = gol2
	}
	if gol3 < golKecil {
		golKecil = gol3
	}
	if gol4 < golKecil {
		golKecil = gol4
	}

	fmt.Print(golBesar, golKecil)
}