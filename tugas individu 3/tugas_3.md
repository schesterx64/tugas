# Tugas Individu: Tipe Data Lanjutan
## 1. Soal Perbandingan Bilangan
```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(a > b, a == b, a < b)
}
```
#### Output
![Output](https://github.com/schesterx64/tugas/blob/main/tugas%20individu%203/perbandinganBilangan_output.png)
## 2. Soal Genap atau Ganjil
```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n % 2 ==0)
}
````
#### Output
![Output](https//github.com/schesterx64/tugas/blob/main/tugas%20individu%203/genapGanjil_output.png)
## 3. Soal Kelipatan Persekutuan
```go
package main
 import "fmt"

 func main () {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	fmt.Println(n % a == 0 && n % b == 0)
 }
```
#### Output
![Output](https://github.com/schesterx64/tugas/blob/main/tugas%20individu%203/kelipatanPersekutuan_output.png)
## 4. Soal Rentang Nilai
```go
package main
 import "fmt"

 func main() {
	var x, low, high int
	fmt.Scan(&x, &low, &high)

	fmt.Println(x >= low && x <= high)
 }
```
#### Output
![Output](https://github.com/schesterx64/tugas/blob/main/tugas%20individu%203/rentangNilai_output.png)