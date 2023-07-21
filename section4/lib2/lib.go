// 라이브러리 접근 제어 (2)
// 사용자 정의 패키지

package lib2 // 폴더 이름

func CheckNum1(c int32) bool { // 소문자가 되면 다른 곳에서 사용 불가
	return c > 100
}

func CheckNum2(c int32) bool { // 소문자가 되면 다른 곳에서 사용 불가
	return c > 1000
}
