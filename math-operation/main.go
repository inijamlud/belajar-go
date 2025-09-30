package main

import "fmt"

func main() {
	x := 10

	x += 2
	x += 8
	x -= 8

	fmt.Println(x)
	fmt.Println(x * 2)
	fmt.Println(x % 2)

	x -= -20
	fmt.Println(x)

	x /= 2
	fmt.Println(x)

	x++
	fmt.Println(x)

}
