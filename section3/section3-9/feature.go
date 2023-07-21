// Go 특징 (1)

package main

import "fmt"

func main() {
	// Go 에서 모호하거나 애매한 문법 금지
	// 후치 연산자 i++ 허용, 전치 전위 연산자 ++i 안됨
	// 증감연산 반환 값 없음 summ := i++ 불가
	// 포인터 허용 연산 비허용
	// 주석 // /* */ 샵 안됨

	// 예제 1
	sum, i := 0, 0
	for i <= 100 {
		// sum += i++ 예러 발생
		sum += i
		i++ // ++i 안됨
	}
	fmt.Println("ex1 : ", sum)

	// 문장 끝 세미콜론 주의
	// 자동으로 문장 끝에 세미 콜론 추가함
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미 콜론 필요 (다만 IDE가 자동으로 갈라줌...)
	// 반복문 for 및 제어문 (조건문) if 사용할 때 주의

	// 예제 2
	for i := 0; i <= 10; i++ { // fmt.Print 는 자동 줄바꿈 x
		// fmt.Print("ex 2 : ", i); fmt.Println(i)
		fmt.Print("ex 2 : ", i)
		fmt.Println(i)
	}
}
