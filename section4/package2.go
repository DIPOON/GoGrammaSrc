// 패키지 접근 제어
package main

import (
	"fmt"
	_ "os" // 사용하지는 않지만 일단 남겨봐 _
	"section4/lib"
	checkUp "section4/lib2" // checkUp 같은 것이 alias
)

func main() {
	var i int32 = 101

	fmt.Println("10 보다 큰 수? : ", lib.CheckNum(i))

	// 변수, 상수, 함수, 메서드, 구조체 등 식별자
	// 대문자 : public. 패키지 외부에서 접근 가능
	// 소문자 : private. 패키지 외부 접근 불가 (패키지 내에서만 접근 가능)

	fmt.Println("100보다 큰 수? : ", checkUp.CheckNum1(i))
	fmt.Println("1000보다 큰 수? : ", checkUp.CheckNum2(i))

	// 별칭 사용
	// 빈 식별자 _ 사용
	fmt.Println(i, "10 보다 큰 수? : ")
}
