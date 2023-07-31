// 데이터 타입 : numeric (1)

package main

import "fmt"

func main() {
	// 데이터 타입 : 숫자형
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수

	// int32 int64 대신 int 추천. 내부적으로 최적화됨
	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex1 : ", num2)
	fmt.Println("ex1 : ", num3)
	fmt.Println("ex1 : ", num4)

	// 정수형 문자 출력
	// 아스키(영문)
	var char1 byte = 72 // int 8 alias
	var char2 byte = 0110
	var char3 byte = 0x48

	fmt.Printf("ex2 : %c %c %c\n", char1, char2, char3) // 문자
	fmt.Printf("ex2 : %d %d %d\n", char1, char2, char3) // digit
	fmt.Printf("ex2 : %d %o %x\n", char1, char2, char3) // 십진수 8진수 16진수

	// 유니코드 (한글)
	var char4 rune = 50556
	var char5 rune = 0142574 // 8진수 44032
	var char6 rune = 0xC57C  // 16진수 44032

	fmt.Printf("ex3 : %c %c %c\n", char4, char5, char6)
	fmt.Printf("ex3 : %d %d %d\n", char4, char5, char6)
	fmt.Printf("ex3 : %d %o %x\n", char4, char5, char6)

	// 실수. 부동소수점
	// float32(7자리), float64(15자리까지 지원)
	var num5 float32 = 0.14
	var num6 float32 = .75647
	var num7 float32 = 442.0378373
	var num8 float32 = 10.0

	fmt.Println("ex4 : ", num5)
	fmt.Println("ex4 : ", num6)
	fmt.Println("ex4 : ", num7)
	fmt.Println("ex4 : ", num8)
	fmt.Println("ex4+ : ", num8-0.1)
	fmt.Println("ex4+ : ", float32(num8-0.1))
	fmt.Println("ex4+ : ", float64(num8-0.1)) // 주의 : 형 변환 중 정확한 값 9.9 를 잃어버렸다.

	// 지수 표기법
	var num9 float32 = 14e6
	var num10 float64 = .156875e+3
	var num11 float64 = 5.32521e-10

	fmt.Println("ex5 : ", num9)
	fmt.Println("ex5 : ", num10)
	fmt.Println("ex5 : ", num11)
}
