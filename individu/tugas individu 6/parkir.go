package main
import "fmt"

func main() {
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1, &h2, &m2)

	if h1 < 7 {
		h1 += 12
	}
	if h2 < 7 {
		h2 += 12
	}

	t1 := (h1 * 60) + m1
	t2 := (h2 * 60) + m2
	selisih := t2 - t1

	hh := selisih / 60
	mm := selisih % 60

	fmt.Printf("%d jam %d menit", hh, mm)
}