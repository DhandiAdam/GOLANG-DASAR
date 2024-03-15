package main

func main() {
	// deklarasi konstan alternatif 1
	const firstName string = "Amir"
	const lastName = "Murtako"

	// error
	firstName = "Ali"
	lastName = "Murtadho"

	// deklarasi konstan alternatif 2
	const (
		namaDepan    = "Desti"
		namaBelakang = "Fitriati"
	)

	// error
	namaDepan = "Ionia"
	namaBelakang = "Veritawati"
}
