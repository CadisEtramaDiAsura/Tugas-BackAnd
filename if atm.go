package main

import "fmt"

func main() {
	var (
		Bahasa   int
		Password int
	)

	fmt.Println("~~~{Selamat Datang}~~ \n \n ")

	fmt.Println("Silahkan Pilih Bahasa \n 1. Bahasa Indonesia \n 2. Bahasa Inggris")

	fmt.Print("Masukkan Bahasa : ")
	fmt.Scan(&Bahasa)
	if Bahasa == 1 {
		fmt.Print("Masukkan Password Anda :")
		fmt.Scan(&Password)
	} else {
		fmt.Println("Maaf Bahasa Yang Anda Masukkan Tidak Ada")
	}
}
