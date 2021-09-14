package main

import "fmt"

func main() {
	//contoh 1
	var daerah1 string
	daerah1 = "Lampung"

	//contoh 2
	var daerah2 int8 = 100

	//contoh 3
	var daerah3 = "Bandung"

	//contoh 4
	daerah4 := "Sumatera utara"

	//contoh 5
	var (
		daerah5   = "Lampung"
		addedWord = "Kode"
	)

	//contoh 1
	fmt.Println(daerah1)

	//contoh 2
	fmt.Println(daerah2)

	//contoh 3
	fmt.Println(daerah3)

	//contoh 4
	fmt.Println(daerah4)

	//contoh 5
	fmt.Println(daerah5, addedWord)

}
