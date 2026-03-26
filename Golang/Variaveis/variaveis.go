package main

import "fmt"

func main() {
	var variavael1 string = "Variavel 1"
	variavel2 := "Varivel 2"

	fmt.Println(variavael1)
	fmt.Println(variavel2)

	var (
		varivael3 string = "lalala"
		varivael4 string = "lalala"
	)
	fmt.Println(varivael3, varivael4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)
}
