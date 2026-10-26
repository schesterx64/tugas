package main
import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n == "A" || n == "I" || n == "U" || n == "E" || n == "O" || n == "a" || n == "i" || n == "u" || n == "e" || n == "o" {
		fmt.Println("Bukan konsonan")
	} else {
		fmt.Println("Konsonan")
	} 
}