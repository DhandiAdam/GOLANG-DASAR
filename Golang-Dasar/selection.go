package main

import "fmt"

func main() {
	name := "Dhandi"

	if name == "Sultan" {
		fmt.Println("Hello Sultan")
	} else if name == "Dhandi" {
		fmt.Println("Hallo Dhandi")
	} else {
		fmt.Println("Halo dek, boleh kenalan ndak?")
	}

	if length := len(name); length > 10 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
