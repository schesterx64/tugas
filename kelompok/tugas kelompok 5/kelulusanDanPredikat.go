package main
import "fmt"

func main() {
	var nilaiUtama, nilaiRemedial int
	fmt.Scan(&nilaiUtama, &nilaiRemedial)

	if nilaiUtama >= 70 {
		fmt.Println("Lulus")
		if nilaiUtama >= 85 {
			fmt.Println("Sangat Baik")
		} else {
			fmt.Println("Baik")
		}
	} else if nilaiUtama < 70 {
		fmt.Println("Remedial")
		if nilaiRemedial >= 70 {
			fmt.Println("Lulus Remedial")
		} else {
			fmt.Println("Tidak Lulus")
		}
	}
}