package main

import "fmt"

func main() {
	fmt.Println("---{Nilai Ganjil Genap}---")

	var nilai int

	fmt.Print("Masukkan Nilai :")
	fmt.Scan(&nilai)

	if nilai%2 == 0 {
		fmt.Println("Nilai : ", nilai, "merupakan bilangan genap")
	} else {
		fmt.Println("Nilai : ", nilai, "merupakan bilangan ganjil")
	}
}
