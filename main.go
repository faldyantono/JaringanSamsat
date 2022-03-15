package main

import (
	"belajargolang/modulego"
	"fmt"
)

func main() {
	//output dari function parameter
	modulego.SayHelloCensored("Haris", modulego.FilterSensor)
	//anonymous struct
	vehicle := struct {
		brand string
		breed string
		class string
		cc    int
	}{
		brand: "Honda",
		breed: "Supra-X",
		class: "Bebek",
		cc:    125,
	}
	fmt.Println("Kendaraan anda adalah              			:", vehicle)
	//output dari return value
	fmt.Println(modulego.Pengguna("Haris"))
	//struct
	pajak := modulego.Pajak{4, 20000000, 2000000}
	fmt.Println("Anda dikenakan Pajak Progresif Sebesar		 	:", pajak.Perkalian(), "Rupiah")
	//anonymous function
	func(i int) {
		fmt.Println("Anda Memiliki Sepeda Motor sebanyak      		:", i, "unit")
	}(4)
	//struct method
	denda := modulego.Denda{2, 3, 0}
	fmt.Println("Anda dikenakan Denda Sebesar				:", denda.Perkalian2(), "Rupiah")
}
