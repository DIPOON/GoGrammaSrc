// 패키지 (1)

package main

// 선언 방법
// import "fmt"
// import "os"

// 요 방식을 더 많이 씀
import (
	"fmt"
	"os"
)

func main() {
	// 패키지는 코드를 구조화하고 재사용합니다.
	// 응집도는 높이고 결합도는 낮추려고
	// Go 또한 마찬가지로 패키지 단위로 독립된/작은 단위의 개발 -> 패키지를 결합해서 코딩 권고
	// 다른 언어도 마찬가지
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 소스 파일들은 디렉토리 명을 패키지 명으로 사용한다.
	// main 함수 파일만 특별히 하나여야함 -> 컴파일러 공유 라이브러리 x 프로그램 시작점 entity point
	// 네이밍 규칙 : 소문자 private 대문자 public

	// 기본 패키지 ex) fmt 는 GOROOT 에 있음
	// 사용자 패키지는 GOPATH

	var name string

	fmt.Println("이름은? : ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
