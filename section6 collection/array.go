// 자료형 : 배열

package main

import "fmt"

func main() {
	// 배열은 용량, 길이가 정해지고 바뀌지 않는다.
	// 배열 vs 슬라이스 중요
	// 배열은 길이 고정 vs 슬라이스는 길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가 section 9~10 참조
	// 슬라이스의 사용이 흔하다

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 개수

	// 예제 1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // 혼자 자료형 다른거면 안됨
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}   // 짧은 선언 가능
	arr5 := [5]int{1, 2, 3}         // 0 들어감
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열 크기 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 마지막 콤마
	} // 2차원 가능

	fmt.Printf("", arr1) // 뭔가 길게 나옴
	fmt.Printf("\n")

	fmt.Printf("arr1 %-5T %d %v\n", arr1, len(arr1), arr1) // integer 기본 초깃값 0
	fmt.Printf("arr2 %-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("arr3 %-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("arr4 %-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("arr5 %-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("arr6 %-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("arr7 %-5T %d %v\n", arr7, len(arr7), arr7)

	arr1[2] = 5
	fmt.Printf("%d %v\n", len(arr1), arr1)

	// 예제 2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}
	fmt.Printf("arr8 %-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("arr9 %-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("arr10 %-5T %d %v\n", arr10, len(arr10), arr10)
	fmt.Println()

	// 예제 3 배열 순회
	arr11 := [5]int{1, 10, 100, 1000, 10000}

	// len 길이 반복
	for i := 0; i < len(arr11); i++ {
		fmt.Println("ex3 : ", arr11[i])
	}
	fmt.Println()

	// 예제 4 많이 쓰이는 순회
	arr12 := [5]int{7, 77, 777, 7777, 77777}
	// range
	for i, v := range arr12 {
		fmt.Println("ex4 : ", i, v)
	}
	// 인덱스 생략1
	for _, v := range arr12 {
		fmt.Println("ex4 : ", v)
	}
	// 생략2 i, _ 와 같다.
	for i := range arr12 {
		fmt.Println("ex4 : ", i)
	}
	fmt.Println()

	// 값을 복사해서 전달한다는 증명
	// 예제 5
	arr13 := [5]int{1, 10, 100, 1000, 10000}
	arr14 := arr13

	fmt.Println("ex5 : ", arr13, &arr13)
	fmt.Println("ex5 : ", arr14, &arr14)

	fmt.Printf("ex5 : %p %v\n", &arr13, arr13) // 포인터, 밸류
	fmt.Printf("ex5 : %p %v\n", &arr14, arr14) // 주소값이 위아래가 다르다.
	arr13[2] = 5
	fmt.Printf("ex5 : %p %v\n", &arr13, arr13)
	fmt.Printf("ex5 : %p %v\n", &arr14, arr14) // 내부 수정 반영 X
	fmt.Printf("ex5 : %p \n", arr13)           // 주소랑 내부가 길게 나옴
}
