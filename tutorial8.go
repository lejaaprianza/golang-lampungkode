package main

import "fmt"

func main() {
	var angka64 int64 = 127
	var angka32 int32 = int32(angka64)
	var angka8 int8 = int8(angka32)

	fmt.Println(angka64)
	fmt.Println(angka32)
	fmt.Println(angka8)
}
