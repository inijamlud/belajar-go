package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name string = "Jamal"
	// var umur int = 20
	// city := "Bandung"
	// year := 2024

	// const PI = 3.14

	// fmt.Println("Name: ", name)
	// fmt.Println("Umur: ", umur)
	// fmt.Println("city: ", city)
	// fmt.Println("year: ", year)
	// fmt.Println("PI: ", PI)

	// var i8 int8 = 127
	// var i16 int16 = 32767

	// var u8 uint8 = 255
	// var u16 uint16 = 65535

	// fmt.Println("Signed Integers")
	// fmt.Printf("int8	: %v\n", i8)
	// fmt.Printf("int16	: %v\n", i16)

	// fmt.Println("\nUnsigned Integers")
	// fmt.Printf("uint8	: %v\n", u8)
	// fmt.Printf("uint16	: %v\n", u16)

	// var f32 float32 = 255.2
	// var f64 float64 = 65535.90

	// fmt.Println("Float")
	// fmt.Printf("f32	: %v\n", f32)
	// fmt.Printf("f64	: %v\n", f64)

	// var _isTrue bool = true
	// var _isFalse bool = false

	// fmt.Println("Boolean")
	// fmt.Printf("true	: %v\n", _isTrue)
	// fmt.Printf("false	: %v\n", _isFalse)

	text := "Halo, Dunia"

	fmt.Println("Lower case: ", strings.ToLower((text)))
	fmt.Println("Upper case: ", strings.ToUpper((text)))
	fmt.Println("Prefix case: ", strings.HasPrefix(text, "Halo"))
	fmt.Println("Contain case: ", strings.Contains(text, "Halo"))

	parts := strings.Split("apel,jagung,parmesan", ",")
	fmt.Println("Split: ", parts)

	newText := strings.ReplaceAll(text, "Dunia", "Malaya")
	fmt.Println("Replace: ", newText)
}
