package main

import "fmt"

func main() {

	//contoh 1
	const daerah1 string = "Lampung"

	//contoh 2
	const daerah2 = "Bandung"

	//contoh 3
	const (
		angka   = 6
		kondisi = true
	)

	//contoh 1
	fmt.Println(daerah1)

	//contoh 2
	fmt.Println(daerah2)

	//contoh 3
	fmt.Println(angka)
	fmt.Println(kondisi)

}
