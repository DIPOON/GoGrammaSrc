package main

import "fmt"

func main() {
	// 짧은 선언
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10

	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	var i int = 10
	if i < 15 {
		fmt.Println("normal variable", i)
	}

	// 짧은 선언이 이미 있는 변수에도 할 수 있네
	if i := 7; i < 11 {
		fmt.Println("short variable", i)
	}

	fmt.Println("variable", i)
}
