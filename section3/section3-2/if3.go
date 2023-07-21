package main

import "fmt"

func main() {
	i := 100

	// if - else if
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100이상 120 미만")
	} else {
		fmt.Println("그 밖에. 아마도 100 미만")
	}
}
