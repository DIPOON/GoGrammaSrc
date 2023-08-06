// 데이터 타입 : Numeric 연산 (1)

package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산 (산술, 비교)
	// 타입이 같아야 가능 : 엄격하다
	// 다른 타입끼리는 형변환 후 연산 가능. 안하면 에러 발생
	// +, -, *, % 나머지, / 몫, <<, >>, &, ^ 보수

	// 예제 1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64 // 최솟값은 Min

	fmt.Println("n1 ", n1)
	fmt.Println("n2 ", n2)
	fmt.Println("n3 ", n3)
	fmt.Println("n4 ", n4)
	fmt.Println("", math.MaxFloat32)
	fmt.Println("", math.MaxFloat64)

	n5 := 100000       // 기본값 int
	n6 := int16(10000) // 형변환 되어 int16이 됨
	// n7 := uint8(300)   // overflow error

	// fmt.Println("n5 + n6 ", n5 + n6) // invalid operation. 다른 타입끼리는 형 맞춰야한다.
	fmt.Println("n5 + n6 ", n5+int(n6)) // int랑 int32 랑 다르다!
	fmt.Println("n5 > int(n6) ", n5 > int(n6))

	var k1 uint8 = 125
	var k2 uint8 = 90

	fmt.Println("k1+k2 ", k1+k2)
	fmt.Println("k1-k2 ", k1-k2)
	fmt.Println("k1*k2 ", k1*k2)
	fmt.Println("k1/k2 ", k1/k2)
	fmt.Println("k1%k2 ", k1%k2)
	fmt.Println("k1<<2 ", k1<<2)
	fmt.Println("k1>>2 ", k1>>2)
	fmt.Println("^k1 ", ^k1)

	var k3 int = 12
	var k4 float32 = 8.2
	var k5 uint16 = 1024
	var k6 uint32 = 120000

	fmt.Println("float32(k3) +k4", float32(k3)+k4)
	fmt.Println("k3+int(k4)", k3+int(k4))         // 값이 다름 : 절삭을 의도한 것이 아니라면 위험
	fmt.Println("k5 + uint16(k6)", k5+uint16(k6)) // 절삭의 또다른 예시

	// 범위 초과
	// var m1 uint8 = math.MaxUint8 + 1 // 오버플로우 에러

	// 범위 미만
	// var m2 uint8 = -1
}
