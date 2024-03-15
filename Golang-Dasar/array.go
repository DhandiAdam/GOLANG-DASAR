package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "Adi Wahyu"
	nama[1] = "Amir Murtako"
	nama[2] = "Bambang Riono"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	values[0] = 100
	fmt.Println(values)
}
