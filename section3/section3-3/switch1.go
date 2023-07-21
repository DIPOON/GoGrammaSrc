// Switch문(1)
package main

import "fmt"

func main() {
	// 제어문(조건문) - switch
	// Switch 뒤 표현식 (expression) 생략 가능
	// case 뒤 표현식 사용 가능
	// 자동 break 때문에 fallthrouth 존재
	// Type 분기 가능 -> 값이 아닌 변수 자료형 Type 으로 분기 가능

	// 예제 1
	a := -7
	switch {
	case a < 0: // s와 c 들여쓰기 같은 정도로
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제 2 go lang 추천 방법
	switch b := 27; {
	case b < 0: // s와 c 들여쓰기 같은 정도로
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제 3 등호 생략하도록 어느 값을 비교하는지
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("No")
	}

	// 예제 4 연산자로 가능!
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("No")
	}

	// 예제 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
