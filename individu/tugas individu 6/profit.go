package main
import "fmt"

func main() {
	var bulanIni, bulanSebelum, selisih float64
	fmt.Scan(&bulanSebelum, &bulanIni)

	if bulanIni > bulanSebelum {
		selisih = bulanIni - bulanSebelum
		fmt.Println("Peningkatan sebesar:", selisih)
	} else if bulanIni < bulanSebelum {
		selisih = bulanSebelum - bulanIni
		fmt.Println("Penurunan sebesar:", selisih)
	} else if bulanIni == bulanSebelum {
		fmt.Println("Tetap")
	}
}