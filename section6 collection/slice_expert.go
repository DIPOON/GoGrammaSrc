// 자료형 : 슬라이스 심화

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 추가 및 병합
	// 예제 1
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)       // 길이 말고 용량을 넘어서 append 하면 재할당 퍼포먼스 낭비 발생. 옹량 두배씩 할당함
	s2 = append(s1, s2...)      // 슬라이스에 슬라이스를 붙여 넣을 때에는 ...
	s3 = append(s2, s3[0:3]...) // s3 추출 후 병합

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	// 예제 2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}

	// 예제 3 슬라이스 추출
	// slice[i:j] i 부터 j-1 까지 추출
	// slice[i:] i 부터 마지막 까지 추출
	// slice[:j] 처음 부터 j-1 까지 추출
	// slice[:] 처음 부터 마지막 까지 추출
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("ex3 : ", slice1[:])
	fmt.Println("ex3 : ", slice1[0:])
	fmt.Println("ex3 : ", slice1[:5])
	fmt.Println("ex3 : ", slice1[0:len(slice1)-1])
	fmt.Println("ex3 : ", slice1[3:])
	fmt.Println("ex3 : ", slice1[:3])
	fmt.Println("ex3 : ", slice1[1:3])

	// 예제 4 정렬
	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice3 := []string{"b", "d", "f", "a", "c", "e"}

	fmt.Println("ex4")
	fmt.Println(sort.IntsAreSorted(slice2))
	sort.Ints(slice2) // 정렬
	fmt.Println(slice2)
	fmt.Println(sort.IntsAreSorted(slice2))

	fmt.Println("ex4")
	fmt.Println(sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println(slice3)
	fmt.Println(sort.StringsAreSorted(slice3))

	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make 로 공간을 할당 하고 나서 복사 해야한다.
	// 복사된 슬라이스 값 변경해도 원본 영향 X

	// 예제 5 복사
	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := make([]int, 5)
	slice6 := []int{}

	copy(slice5, slice4) // 5개만 복사되고 짤림
	copy(slice6, slice4) // 빈 슬라이스

	fmt.Println("ex5 : ", slice5)
	fmt.Println("ex5 : ", slice6)

	// 예제 6
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex6 : ", a)
	fmt.Println("ex6 : ", b)

	// 예제 7 copy 주의
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 부분적으로 슬라이스 추출하면 참조
	d[1] = 7
	fmt.Println("ex7 : ", c) // d를 바꿨는데 c도 바뀌었다!
	fmt.Println("ex7 : ", d)

	// 예제 8 용량 지정 copy
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 7은 용량
	fmt.Println("ex8")
	fmt.Printf("ex8 -> len : %d, cap : %d, value : %v\n", len(f), cap(f), f)
}
