// Go 초기화 함수

package main

import (
	"fmt"
	"section4/lib2" // 여기의 init 부터 실행된다.
)

var num int32 // 완전 전역 변수

func init() {
	num = 300
	fmt.Println("Init Method Start!")
}

func main() {
	// init : 패키지 로드 시에 가장 먼저 호출
	// 가장 먼저 초기화되는 작업 작성 시 유용하다.
	fmt.Println("100보다 큰 수? ", lib2.CheckNum1(num))
}
