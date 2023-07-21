// Go 특징 (2)

package main

import "fmt"

func main() {
	// 코드 서식 지정
	// 코딩 스타일 유지
	// gofmt -h 사용법
	// gofmt -w 원본 파일에 반영

	// 예제1
	// for i:= 0; i<= 100; i++{
	// 	fmt.Println("ex1 : ", i)
	// }
	/* 이 주석 내부의 코드도 수정해버리네;;; */

	for i := 0; i <= 100; i++ {
		fmt.Println("ex1 : ", i)
	}
}
