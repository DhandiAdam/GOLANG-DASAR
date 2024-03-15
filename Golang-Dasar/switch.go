package main

import "fmt"

func main() {
	name := "Euis"

	switch name {
	case "Euis":
		fmt.Println("Hello Euis")
	case "Dina":
		fmt.Println("Hello Dina")
	default:
		fmt.Println("Halo dek, boleh kenalan ndak?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
