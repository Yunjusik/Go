//패키지 접근제어 (1)
package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	//패키지 접근제어
	//변수, 상수, 함수, 메서드, 구조체 등 식별자
	//대문자 : 첫글자패키지 외부에서 접근 가능
	//소문자 : 패키지 외부 접근 불가 (해당 패키지내에서만 접근 가능한 private)
	fmt.Println("100보다 큰 수인가? : ", lib2.CheckNumb1(101))
	fmt.Println("1000보다 큰 수인가? : ", lib2.CheckNumb2(888))
}
