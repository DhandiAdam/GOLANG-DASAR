package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// function parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function with return value
func getHello(name string) string {
	return "Hello " + name
}

// function with multiple values
func getFullName() (string, string) {
	return "Supra", "Hatmara"
}

// function with named return values
func getCompleteName() (namaDepan, namaTengah, namaBelakang string) {
	namaDepan = "Ahmad"
	namaTengah = "Dhani"
	namaBelakang = "Prasetyo"

	return namaDepan, namaTengah, namaBelakang
}

// Fungsi hitungLuasKeliling dengan named return values
func hitungLuasKeliling(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	// Return multiple values with named return values
	return
}

// Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	sayHello()

	sayHelloTo("Tengku", "Firmansyah")

	siapaya := getHello("Dinda")
	fmt.Println(siapaya)

	firstName, lastName := getFullName()
	fmt.Println("Hello", firstName, lastName)

	// menghiraukan Return Value
	namaPertama, _ := getFullName()
	fmt.Println("Hello", namaPertama)

	// named return values
	namaDepan, namaTengah, namaBelakang := getCompleteName()
	fmt.Println("Hello", namaDepan, namaTengah, namaBelakang)

	panjang := 5
	lebar := 3

	fmt.Println("Panjang:", panjang)
	fmt.Println("Lebar:", lebar)

	// memanggil fungsi hitungLuasKeliling dengan paramter panjang dan lebar
	luas, keliling := hitungLuasKeliling(panjang, lebar)
	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)

	// memanggil variadic function
	total := sumAll(10, 10, 10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// slice
	numbers := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
