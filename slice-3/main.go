package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 5)
	fmt.Println(s)

	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	s = append(s, 10, 20)
	fmt.Println("append: ", s)

	slice2 := make([]int, 3)
	slice3 := copy(slice2, s)
	fmt.Println("slice 3: ", slice2, slice3)

	angka := []int{1, 2, 3, 4, 5}
	slice4 := angka[1:3]

	fmt.Println("slice 4: ", slice4)

}
