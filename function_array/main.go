package main

import "fmt"

func main() {
	number := [5]int{1, 2, 3, 4, 5}

	fmt.Println("jumlah elemen: ", len(number))
	fmt.Println("index ke-1: ", number[1])
	number[1] = 89
	fmt.Println("index ke-1: ", number[1])

	fmt.Println("ini adalah array (after)", number)

	for i := 0; i < len(number); i++ {
		fmt.Println("loop", i)

	}

	for idx, v := range number {
		fmt.Println(`loop `, idx, v)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	fmt.Println(`arr1 == arr2 `, arr1 == arr2)
	fmt.Println(`arr1 != arr2 `, arr1 != arr2)

}
