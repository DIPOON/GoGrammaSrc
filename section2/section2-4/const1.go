package main

import "fmt"

func main() {
	// var 대신 const. 선언과 동시에 초기화
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight() 겟 헤잇의 함수가 항상 같은 값을 return한다는 보장이 없으니 에러
	const e = 35.6
	const f = false

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	// fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
}
