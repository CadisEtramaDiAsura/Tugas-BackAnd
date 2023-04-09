package main

import "fmt"

func main() {
	fmt.Println("===[Nilai KKM]===")

	var nilai int

	fmt.Print("Masukkan Nilai :")
	fmt.Scan(&nilai)

	if nilai >= 75 {
		fmt.Println("Nilai anda diatas KKM yaitu ", nilai, " Selamat Anda Lulus ")
	} else {
		fmt.Println("Nilai anda dibawah KKM yaitu ", nilai, " Belajar yang giat yah ")
	}
}
