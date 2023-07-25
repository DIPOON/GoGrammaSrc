// 데이터 타입 : Bool
package main

import "fmt"

func main() {
	// Bool. Boolean : 참과 거짓
	// 조건부 논리 연산자와 주로 사용 : ! (not), || (or), && (and)
	// 암묵적 형 변환 X : 0, Nil -> false 안됨

	// 예제 1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b3)
	fmt.Println("ex5 : ", b1 && b3)
	fmt.Println("ex6 : ", b1 || b2)
	fmt.Println("ex7 : ", !b1)

	// b4 가 명시적으로 true 나 false 가 아니면 형 변환 x
	// b4 := 1
	// if b4 {
	// 	fmt.Println("ex8 : true!")
	// }

	// 예제 2
	fmt.Println("Redundant start")
	fmt.Println("ex : ", true && true)
	fmt.Println("ex : ", true && false)
	fmt.Println("ex : ", false && false)
	fmt.Println("ex : ", true || true)
	fmt.Println("ex : ", true || false)
	fmt.Println("ex : ", false || false)

	// 예제 3
	num1 := 15
	num2 := 37
	fmt.Println("ex3 Start")
	fmt.Println("ex : ", num1 < num2)
	fmt.Println("ex : ", num1 > num2)
	fmt.Println("ex : ", num1 >= num2)
	fmt.Println("ex : ", num1 <= num2)
	fmt.Println("ex : ", num1 == num2)
	fmt.Println("ex : ", num1 != num2)
}
