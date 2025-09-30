package main

import "fmt"

func main() {
	// if v := 20; v > 10 {
	// 	fmt.Println("Hasil lebih dari 10: ", v)
	// }

	var nilai int
	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scan(&nilai)

	if nilai < 0 {
		fmt.Printf("Angka kurang dari 0-100")
	} else if nilai < 60 {
		fmt.Printf("Nilai %d = D", nilai)
	} else if nilai < 70 {
		fmt.Printf("Nilai %d = C", nilai)
	} else if nilai < 85 {
		fmt.Printf("Nilai %d = B", nilai)
	} else if nilai < 90 {
		fmt.Printf("Nilai %d = A-", nilai)
	} else if nilai <= 100 {
		fmt.Printf("Nilai %d = A", nilai)
	} else if nilai > 100 {
		fmt.Printf("Angka melebihi 0-100")

	}
}
