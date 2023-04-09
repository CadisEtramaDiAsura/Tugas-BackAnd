package main

import "fmt"

func main() {
	fmt.Println(" \n \n ")
	fmt.Println("===menghitung luas lingkarang===")
	const phi float32 = 3.14
	var JariJari float32

	JariJari = 12.0

	LuasLingkarang := phi * JariJari * JariJari

	fmt.Println("jika phi =", phi, "dan \n Jari-Jari =", JariJari)
	fmt.Println("Maka Luas Lingkaranya Adalah ", LuasLingkarang)

	fmt.Println(" \n \n ")
	fmt.Println("---menghitung luas segitiga---")
	var (
		alas   float32
		tinggi float32
	)
	alas = 12.4
	tinggi = 10.1
	LuasSegitiga := alas * tinggi * 1 / 2

	fmt.Println("jika Alas =", alas, "dan \n Tinggi =", tinggi)
	fmt.Println("Maka Luas Segitiganya Adalah ", LuasSegitiga)

	fmt.Println(" \n \n ")
	fmt.Println("~~~menghitung luas kubus~~~")
	sisi := 4
	LuasKubus := sisi * sisi

	fmt.Println("jika sisi sebuah kubus adalah", sisi)
	fmt.Println("Maka Luas Kubus Adalah ", LuasKubus)

}
