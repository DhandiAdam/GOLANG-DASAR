package main

import "fmt"

func main() {
	// aritmatika
	var a = 10
	b := 10
	c := a + b
	fmt.Println(c)

	i := 10
	i += 10
	fmt.Println(i)

	j := 1
	j++
	j++
	fmt.Println(j)

	// Boolean untuk kasus string
	nama1 := "Amir"
	nama2 := "Amir"

	var hasil bool = nama1 == nama2

	fmt.Println(hasil)

	// Boolean untuk kasus perbandingan
	nilaiAkhir := 90
	absensi := 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	fmt.Println(lulusNilaiAkhir) // true
	fmt.Println(lulusAbsensi)    // false

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) // false
}
