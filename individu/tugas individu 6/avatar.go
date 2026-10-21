package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)
	appaDewasa := 3
	appaKecil := 5

	dewasa := (p + 4) / 5
	if dewasa > appaDewasa {
		dewasa = appaDewasa
	}
	sisaOrang := p - (dewasa * 5)

	kecil := 0
	if sisaOrang > 0 {
		kecil = (sisaOrang + 1) / 2
		if kecil > appaKecil {
			kecil = appaKecil
		}
	}
	sisaOrang -= (kecil * 2)

	takBerangkat := 0
	if sisaOrang > 0 {
		takBerangkat = sisaOrang
	}

	switch {
	case dewasa > 0 && kecil > 0 && takBerangkat > 0:
		fmt.Printf("Dewasa %d, kecil %d, dan %d tak berangkat", dewasa, kecil, takBerangkat)
	case dewasa > 0 && kecil > 0:
		fmt.Printf("Dewasa %d, kecil %d", dewasa, kecil)
	case dewasa > 0 && takBerangkat > 0:
		fmt.Printf("Dewasa %d, dan %d tak berangkat", dewasa, takBerangkat)
	case dewasa > 0:
		fmt.Printf("Dewasa %d", dewasa)
	case kecil > 0 && takBerangkat > 0:
		fmt.Printf("Kecil %d, dan %d tak berangkat", kecil, takBerangkat)
	case kecil > 0:
		fmt.Printf("Kecil %d", kecil)
	case takBerangkat > 0:
		fmt.Printf("%d tak berangkat", takBerangkat)
	}
}