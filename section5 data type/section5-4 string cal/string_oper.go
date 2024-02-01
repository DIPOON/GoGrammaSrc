// 데이터 타입 : 문자열 연산

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제1 추출. 문자열은 배열
	var str1 string = "Golang"
	var str2 string = "World"
	fmt.Println("ex1-1 : ", str1[0:2], str1[0]) // (파이썬표현)슬라이싱 -> 문자. 그냥 하나 가져오면 숫자
	fmt.Println("ex1-2 : ", str2[3:])           // 비워놓으면 끝까지
	fmt.Println("ex1-3 : ", str2[:4])
	fmt.Println("ex1-4 : ", str1[1:3])

	// 예제2 비교
	fmt.Println("ex2-1 : ", str1 == str2)
	fmt.Println("ex2-2 : ", str1 != str2)
	fmt.Println("ex2-3 : ", str1 > str2) // 주의 : 문자열 비교는 아스키 코드 사전식 비교
	fmt.Println("ex2-4 : ", str1 < str2)

	// 예제3 결합 : 일반 연산
	str3 := "The Go programming language is an open source project to make programmers more productive" +
		"Go is expressive, concise, clean, and efficient" +
		" It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language"
	str4 := "Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection"
	fmt.Println("ex3-1 : ", str3+str4) // 새로 계속 문자열을 생성하는 비효율 발생

	// 예제4 결합 : Join 추천
	strSet := []string{} // 자료형 슬라이스 선언
	strSet = append(strSet, str3)
	strSet = append(strSet, str4)
	fmt.Println("ex4-1 : ", strings.Join(strSet, "-----"))
}
