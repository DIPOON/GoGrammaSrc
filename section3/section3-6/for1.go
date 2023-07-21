// For문 (1)

package main

import "fmt"

func main() {
	// 반복문 - For
	// Go에서 유일하게 제공되는 반복문 ... 확인 필요할듯
	// 다양한 사용법 숙지

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// 에러 발생 1
	/*
		for i := 0; i < 5; i++ // 여기서 세미 콜론을 찍어버림
		{ // 이렇게 띄어서 하면 안됨, 중괄호 생략해도 안됨

		}
	*/

	// 예제 2 무한 루프, while(true)
	// for { // 조건식 초기화 증가식 전부 생략
	// 	fmt.Println("ex2 : Hello, Golang!")
	// 	fmt.Println("ex2 : Hello, Golang!")
	// }

	// 예제 3 Range 용법
	loc := []string{"Seoul", "Busan", "Incheon"} // 서울 인덱스 0, 부산 인덱스 1
	for index, name := range loc {               // 첫번째 것이 인덱스가 되고 두번째 것이 값이됨
		fmt.Println("ex3 : ", index, name)
	}
	for _, name := range loc { // 묵음
		fmt.Println("ex3 : ", name)
	}
	for index := range loc { // 첫번째 것은 인덱스
		fmt.Println("ex3 : ", index)
	}
}
