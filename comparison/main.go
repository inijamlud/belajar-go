package main

import (
	"fmt"
)

func main() {
	var angka1, angka2 int
	// input
	fmt.Printf("angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Printf("angka kedua: ")
	fmt.Scanln(&angka2)

	fmt.Println("hasil banding")
	fmt.Printf("%d == %d ? %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1, angka2, angka1 <= angka2)

}
