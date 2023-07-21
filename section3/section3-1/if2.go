package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 이하")
	}

	if b >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 이하")
	}

	// else 문도 여는 괄호 다음줄 x
}
