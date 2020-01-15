//패키지 접근제어(2)
package main

import (
	"fmt"
	checkUp "section4/lib" //alias를 줘서 긴 import문을 축약어 처럼 쓸 수 있음. 마치  numpy를 np라 하듯이.
	_ "section4/lib2"      // alias자리에 _ 하면 지우진 않고 내비둠. (Go에서는 안쓰는 패키지는 run시 다 삭제됨)
)

func main() {
	// 패키지 접근제어
	// 별칭사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNumb(11))

}
