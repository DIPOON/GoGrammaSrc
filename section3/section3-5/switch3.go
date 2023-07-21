// Switch문(3)
package main

import (
	"fmt"
)

func main() {
	// 예제 1 : 조건을 나열해서 걸 수 있다.
	a := 30 / 15
	switch a {
	case 2, 4, 6: // a가 2, 4, 6인 경우
		fmt.Println(a, " a는 짝수")
	case 1, 3, 5:
		if a >= 20 {
			fmt.Println("switch 내부에서 조건을 걸 수 있다")
		}
		fmt.Println(a, " a는 홀수")
	}

	// 예약어 fallthrough 존재 의미
	// 예제 2
	// 지금까지 break 없이 그냥 넘어감
	// fallthrough 로 조건 무시하고 아래 처리
	switch e := "go"; e {
	case "java":
		fmt.Println("Java") // 여기는 실행 안됨
		fallthrough
	case "go":
		fmt.Println("Go")
		fallthrough
	case "python":
		fmt.Println("Python")
		fallthrough
	case "ruby":
		fmt.Println("Ruby")
		fallthrough
	case "c":
		fmt.Println("C")
		// fallthrough : cannot fallthrough final
	}
}
