// Go 초기화 함수

package main

import (
	"fmt"
	_ "section4/lib2" // 여기의 init 부터 실행된다.
)

func init() {
	fmt.Println("Init1 Method Start!")
}

func init() {
	fmt.Println("Init2 Method Start!")
}

func main() {
	// init : 패키지 로드 시에 가장 먼저 호출
	// 가장 먼저 초기화되는 작업 작성 시 유용하다.
	fmt.Println("Main Method Start!")
}
