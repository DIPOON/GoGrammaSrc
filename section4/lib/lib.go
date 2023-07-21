// 파일명은 상관 없다.

// 라이브러리 접근 제어 (1)
// 사용자 정의 패키지

package lib // 폴더 이름

func CheckNum(c int32) bool { // 소문자가 되면 다른 곳에서 사용 불가
	return c > 10
}
