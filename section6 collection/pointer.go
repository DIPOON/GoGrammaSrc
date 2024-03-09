// 자료형 포인터

package main

import "fmt"

func main() {
	// Go 포인터 지원
	// 변수의 지역성때문에 사용, 연속된 메모리 참조, 힙 스택
	// 파이썬, 자바 포인터, C# 지원 X. 내부 JRE, 컴파일러, 인터프리터 등은 활용 ... C#이 지원 안한다고???
	// 주소의 값은 직접 변경 불가능. 잘못된 코딩으로 인한 버그 방지
	// * 애스터리스크로 사용
	// nil 로 초기화 (nil == 0)

	// 예제 1
	var a *int            // 방법 1
	var b *int = new(int) // 방법 2

	fmt.Println("ex 1 :")
	fmt.Println(a) // nil
	fmt.Println(b) // 주소값

	// a = 5 에러. &로된 값만 전달 가능. a++ a = 0x981375 이런거 다 막아놓음

	i := 7
	fmt.Println(i, &i)

	a = &i // & 주소값 전달
	b = &i

	fmt.Println("a setting")
	fmt.Println(a, &i)
	fmt.Println(&a)
	fmt.Println(*a) // 역참조
	fmt.Println("b setting")
	fmt.Println(b, &i)
	fmt.Println(&b)
	fmt.Println(*b)

	var c = &i
	d := &i
	*d = 10 // 역참조라서 값이 바뀐다
	fmt.Println("c setting")
	fmt.Println(c, &i)
	fmt.Println(&c)
	fmt.Println(*c)
	fmt.Println("d setting")
	fmt.Println(d, &i)
	fmt.Println(&d)
	fmt.Println(*d)

	// 예제 2
	fmt.Println("ex2")

	ex2i := 7
	ex2p := &ex2i

	fmt.Println(ex2i, *ex2p, &ex2i, ex2p)

	*ex2p++ // 1 증가
	fmt.Println(ex2i, *ex2p, &ex2i, ex2p)

	*ex2p = 7777 // 포인터 변수 역 참조 값 변경
	fmt.Println(ex2i, *ex2p, &ex2i, ex2p)

	ex2i = 88
	fmt.Println(ex2i, *ex2p, &ex2i, ex2p)

	// 예제 3 포인터 값 전달
	// 함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메서드 내에서는 원본 값 변경 불가능
	// 원본 값 변경 위해서 포인터로 전달
	// 특히 크기가 큰 배열인 경우 값 복사 시에 연산 부담 -> 포인터 전달로 해결 (슬라이스, 맵은 참조 전달이라 필요 X)
	fmt.Println("ex3")
	var ex3a int = 10
	var ex3b int = 10
	fmt.Println(ex3a, ex3b)

	rptc(&ex3a) // 값 바뀜
	vptc(ex3b)  // 값이 안바뀌었다.
	// vptc(&ex3b) // 에러발생
	fmt.Println(ex3a, ex3b)

}

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}
