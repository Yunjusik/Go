//Go 특징3
package main

import "fmt"

func main() {
	//코드서식을 지정해주는 유틸리티
	//한 사람이 코딩한 것 같은 일관성 유지
	//코드  style 유지
	// gofmt -h : 사용법 알려줌
	//gofmt -w xxx.go: 원본파일에 반영

	//예제1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex1 : ", i)
  }

}
