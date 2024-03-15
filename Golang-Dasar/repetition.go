package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for konter := 1; konter <= 10; konter++ {
		fmt.Println("Perulangan ke-", konter)
	}

	// array
	names := []string{"Fahri", "Darmawan", "Suryapratama", "Anggraini"}

	// for range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke-", i)
	}

	// continue
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke-", x)
	}
}
