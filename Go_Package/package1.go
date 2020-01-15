// 패키지 (1)
package main
//선언 방법1
//import "fmt"
//import :"os"
//선언 방법2
import (
  "fmt"
  "os"
)
func main() {
  // 패키지 : 코드 모듈화 및 재사용. 라이브러리화
  // go : 패키지단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램 작성 권고
  // 패키지 이름 = 디렉토리 이름
  // 같은 패키지 내 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
  // 네이밍 규칙 : 소문자 private, 대문자: public ..
  // go:main package는 특별하게 인식됨... -> 컴파일러 공유 라이브러리 X , 프로그램의 시작점 start point

    var name string
    fmt.Println("name :" )
    fmt.Scanf("%s", &name)
    fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

}
