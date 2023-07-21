// 열거형3
package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)

	fmt.Println(A)
	fmt.Println(C)

	const (
		_ = iota + 0.75*2
		DEFAULT // 1 + 1.5 = 2.5
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
