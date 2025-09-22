package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 5, 6}

	s1 := arr[:]
	s2 := arr[:3]
	s3 := arr[2:]
	s4 := arr[4:5]
	fmt.Println(s1, s2, s3, s4)
}
