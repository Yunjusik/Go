//go 초기화함수(1)
package main

import (
	"fmt"
)

func init() { //main보다 먼저 실행됨
	//init : 패키지 로드시 가장 먼저 호출
	//가장 먼저 초기화되는 작업 작성 시 유용
	fmt.Println("Init Method Start!")
}

func main() {
	fmt.Println("Main method start!")
}
