package main

import "fmt"

func main() {
	names := [...]string{"Dhandi", "Adam", "Daffa Fathan", "Naupal", "Dosen Adi", "Dosen Sri", "Dosen Desti", "Dosen Bambang"}
	slice := names[1:6] // slice dari elemen index 2 hingga 5

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	// append
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	daySlice1[2] = "Senin Baru"
	fmt.Println(days)

	daySlace2 := append(daySlice1, "Libur baru")
	daySlace2[0] = "Ups"
	fmt.Println(daySlace2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dhandi"
	newSlice[1] = "Adam"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println("fromSlice = ", fromSlice)
	fmt.Println("toSlice = ", toSlice)

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
