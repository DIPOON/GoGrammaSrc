// For문 (2)

package main

import "fmt"

func main() {
	// 예제 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 = sum1 + i // sum1 += i 도 됨
	}
	fmt.Println("ex1 : ", sum1)

	// 예제 2
	sum2, i := 0, 0
	for i <= 30 {
		sum2 += i
		i++
		// j := i++ 에러. 후치 연산 반환값 x
	}
	fmt.Println("ex2 : ", sum2)

	// 예제 3 while 과 유사
	sum3, i := 0, 0
	for {
		if i > 40 {
			break // 가장 가까운 반복문 탈출
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	// 예제 4
	for i, j := 0, 0; i <= 3; i, j = i+1, j+10 { // 초기화 이후에는 i := 이렇게 짧은 정의 또 하면 안됨
		fmt.Println("ex4 : ", i, j)
	}

	// 에러 발생
	// for i, j := 0, 0; i <= 3; i++, j+= 10 {
	// 	fmt.Println("ex4 : ", i, j)
	// }
}
