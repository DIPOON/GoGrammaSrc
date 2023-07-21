// For문 (3)

package main

import "fmt"

func main() {
	// 예제 1 break 레이블 주기
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	// 예제 2 continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 : ", i)
	}

Loop2:
	// Loop 레이블 밑에 관련 없는 소스코드가 나오면 에러
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2 // 1, 2 일때 생략하고 i문 for 문으로 다시
			}
			fmt.Println("ex3 : ", i, j)
		}
	}
}
