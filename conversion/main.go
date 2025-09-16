package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var a int = 10
	// var b float64 = float64(a)

	// var scoreText string = strconv.Itoa(a)

	// fmt.Println(a, b, scoreText)

	// var text string = "100A"          // 100
	// number, err := strconv.Atoi(text) // string to int

	// if err != nil {
	// 	fmt.Println("err:", err.Error())
	// } else {
	// 	fmt.Println("number:", number)
	// }
	// fmt.Println("nil:", nil)

	// boolean to string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println(str)

	val, _ := strconv.ParseBool("true")
	fmt.Println("strong ke bool", val)

}
