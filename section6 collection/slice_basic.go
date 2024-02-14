// 자료형 : 슬라이스 기초

package main

import "fmt"

func main() {
	// 길이 가변. 동적으로 크기가 늘어난다. 레퍼런스 (참조) 타입
	// 슬라이스에는 길이와 용량 개념이 있음
	// 배열 vs 슬라이스 참고 (array.go)
	// 선언 방법 2가지 1. 배열 처럼 선언, 2. make 함수 : make(자료형, 길이, 용량(디폴트 길이))

	// 예제 1
	var slice1 []int // 사이즈 지정 안하면 자동으로 slice. [...] 안됨
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	slice3[4] = 10 // 값 수정 가능

	fmt.Printf("slice1 %-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 %-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("slice3 %-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("slice4 %-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// slice2[0] = 1 안됨

	// 예제 2 make
	var slice5 []int = make([]int, 5, 10) // 길이는 5지만 용량 10. 용량 재할당이 처리량이 있기에 미리 잡아놓을 수 있음
	var slice6 = make([]int, 5)           // 흔한 사용법
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7                                                                  // 삽입
	fmt.Printf("slice5 %-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5) // make 에서는 길이만큼 0 초기화된다.
	fmt.Printf("slice6 %-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("slice7 %-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("slice8 %-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	// 예제 3
	var slice9 []int // nil 길이 용량 0 슬라이스
	if slice9 == nil {
		fmt.Println("This is Nil Slice!")
	}

	// 슬라이스 참조 타입 확인
	// 예제 4 배열 복습
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1 // 복사
	arr2[0] = 7

	fmt.Println("ex4 : ", arr1)
	fmt.Println("ex4 : ", arr2)
	fmt.Println()

	// 예제 5 슬라이스
	slice10 := []int{1, 2, 3}
	var slice11 []int
	slice11 = slice10
	slice11[0] = 7
	fmt.Println("ex5 : ", slice10)
	fmt.Println("ex5 : ", slice11)

	// 예제 6 슬라이스 예외 상황
	slice12 := make([]int, 50, 100)
	fmt.Println("ex6 : ", slice12[4])
	// fmt.Println("ex6 : ", slice12[50]) // 길이 초과 panic
	fmt.Println()

	// 예제 7 슬라이스 순회
	for i, v := range slice10 {
		fmt.Println("ex4 : ", i, v)
	}
}
