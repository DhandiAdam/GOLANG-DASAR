package main

import "fmt"

func main() {
	// // alternatif pertama
	// var name string

	// name = "Adi Wahyu Pribadi"

	// // alternatif kedua
	// var name = "Adi Wahyu Pribadi"

	// // alternatif ketiga
	name := "Adi Wahyu Pribadi"
	fmt.Println(name)

	name = "Adi Wahyu"
	fmt.Println(name)

	// deklarasi multi variabel
	var (
		firstName = "Amir"
		lastName  = "Murtako"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
