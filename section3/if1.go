// If 문 첫번째 파일

package main

import "fmt"

func main() {
	var a int = 20
	b := 20
	if a >= 15 {
		fmt.Println("15 이상이다")
	}

	if b > 25 {
		fmt.Println("25 이상이다")
	}

	// 자주 발생하는 에러 헉... 괄호 여는 것이 if 와 같은 줄에 있어야한다.
	// 자주 발생하는 에러 헉... 괄호 생략하면 안됨
	// 1, 0 자동 형변환 안됨.

	if c := 40; c >= 35 {
		fmt.Println("35 이상이다")
	}
	// 여기서 c += 20 안됨.
}
