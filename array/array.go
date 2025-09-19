package main

import "fmt"

func main() {
	var number [3]int = [3]int{1, 2, 3}
	fmt.Println(number)
	fmt.Println(number[2], number[0])

	number[1] = 80
	fmt.Println(number)
}
