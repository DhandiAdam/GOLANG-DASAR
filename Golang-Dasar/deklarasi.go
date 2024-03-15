package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAdi NoKTP = "123456789"

	fmt.Println(ktpAdi)
	fmt.Println(NoKTP("987654321"))
}
