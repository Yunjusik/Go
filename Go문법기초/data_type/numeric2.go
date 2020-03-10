// 데이터 타입  :숫자2

package main

import "fmt"

func main() {
	// 데이터타입 : 숫자형일
	//정수형 문자 출력
	//예제1
	// 아스키(영문)
	// byte타입은 uint8과 같음

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	//예제2
	//유니코드(한글)
	var char4 rune = 50556   //유니코드
	var char5 rune = 0142574 //44032 (8진수)
	var char6 rune = 0xc57c  //44032(16진수)

	//보통 byte랑 rune을 많이쓰는듯
	fmt.Printf("%c %c %c\n", char1, char2, char3) //아스키코드로반환
	fmt.Printf("%d %d %d\n", char1, char2, char3) //숫자로반환
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6) //유니코드
	fmt.Printf("%d %d %d\n", char4, char5, char6) //10진수
	fmt.Printf("%d %o %x\n", char4, char5, char6) //10진수  8진수 16진수
}

//%c, %d, %o, %x
