// 자료형 맵

package main

import "fmt"

func main() {
	// 맵 (Map)
	// 해시테이블, 딕셔너리(파이썬)
	// Key-Value로 자료 저장
	// 레퍼런스 타입 (참조값 전달)
	// 비교 연산자 사용 불가능 (참조 타입이므로)
	// 참조 타입으로는 Key에 사용 불가. Value로는 모든 타입 가능
	// make 함수 및 축약 (리터럴) 로 초기화 가능
	// 순서 없음

	// 예제 1
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 리터럴형

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	// 예제 2
	map4 := map[string]int{} // Json 형태
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33
	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}
	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])

	// 예제 3 맵 조회 및 순회 (Iterator)
	map11 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}
	fmt.Println("ex3 : ", map11["google"])
	fmt.Println("ex3 : ", map11["daum"])
	fmt.Println()

	// 예제 4 순서 없으므로 랜덤
	for k, v := range map11 {
		fmt.Println("ex4 : ", k, v)
	}
	for k, v := range map11 { // 실행 중에도 순서 보장 X
		fmt.Println("ex4 : ", k, v)
	}
	for _, v := range map11 { // 스킵 가능
		fmt.Println("ex4 : ", v)
	}
	fmt.Println()

	// 예제 5 맵 값 변경 및 삭제
	fmt.Println("ex5 : ", map11)
	map11["home"] = "localhost" // 없을 때 추가
	fmt.Println("ex5 : ", map11)
	map11["home"] = "http://test.com" // 있을 때 수정
	fmt.Println("ex5 : ", map11)
	delete(map11, "home") // 삭제
	fmt.Println("ex5 : ", map11)

	// 예제 6 맵 조회할 경우의 주의 및 팁
	map12 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}
	value1 := map12["lemon"]
	value2 := map12["kiwi"]       // 존재 X. 변수형 기본값 int 0, string "", float 0.0
	value3, ok1 := map12["kiwi"]  // 키값 존재를 확인하기 위해 true, false
	value4, ok2 := map12["lemon"] // 두번째 리턴값으로 키값 존재 확인

	fmt.Println("ex6 : ", value1)
	fmt.Println("ex6 : ", value2)
	fmt.Println("ex6 : ", value3, ok1)
	fmt.Println("ex6 : ", value4, ok2)

	// 예제 7
	if value, ok := map12["kiwi"]; ok {
		fmt.Println("ex7 : ", value)
	} else {
		fmt.Println("ex7 : kiwi is not exist!")
	}
	if value, ok := map12["banana"]; ok {
		fmt.Println("ex7 : ", value)
	} else {
		fmt.Println("ex7 : banana is not exist!")
	}
	// if _, ok := map12["kiwi"]; !ok { // _ 사용 가능, 키값 없을 때 예외 처리할 때 많이 사용
	// 	fmt.Println("ex7 : kiwi is not exist!")
	// } // ??? 윈도우 디펜더에서 바이러스 있어보인다고 막네요 ???
}
