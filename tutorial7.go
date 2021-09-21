package main

import "fmt"

func main() {

	//contoh 1
	var daerah1 string
	daerah1 = "Lampung"
	//tidak bisa :
	// const daerah1 string
	// daerah1 = "Lampung"

	//contoh 2
	const daerah2 = "Bandung"

	//contoh 1
	fmt.Println(daerah1)

}
