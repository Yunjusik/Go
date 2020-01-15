//go 초기화함수(1)
package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init1 Method Start!")
}

// main func없으면 실행 안되고, 패키지형식으로만 됨 (ex: lib)
//패키지형식 파일은 단독실행이 불가능하며, 함수들만 정의해놓고 씀
//만약 패키지 라이브러리용 파일에 init method를 사용한다면, 본 파일에 패키지를 불러올때마다 init이 호출됨
func init() {
	fmt.Println("Init2 Method Start!")
}
func init() {
	fmt.Println("Init3 Method Start!")
} // init method 갯수는 상관없지만 굳이 여러개로 선언할 이유는 없음
func main() {
	fmt.Println("Main Method Start!")
}
