//switch문 (1)
package main
import "fmt"
func main(){
  // 제어문(조건문) - switch
  // switch뒤 표현식 생략 가능
  // case 뒤 표현식 (Expression) 사용 가능
  // 자동 break로 인한 fallthrouth 존재
  // TYpe 분기 가능 -> 값이 아닌 변수타입으로 분기 가능

  //예제 1
  a := -7
  switch {
  case a < 0:
    fmt.Println(a,"는 음수")
  case a ==0:
    fmt.Println(a,"는 0")
  case a > 0 :
    fmt.Println(a,"는 양수")
  }
//예제2
switch b := 27; {
  case b < 0:
    fmt.Println(b,"는 음수")
  case b ==0:
    fmt.Println(b,"는 0")
  case b > 0 :
    fmt.Println(b,"는 양수")
  }
//에제3
switch c := "go"; c { //c를 위로 올려서 각 case에 expression을 생략해버렸다.
case "go":
  fmt.Println("Go!")
case "java":
  fmt.Println("Java!")
default:
  fmt.Println("일치하는 값 없음")
}

//예제 4
 switch c:= "go"; c + "lang" {
   case "golang":
     fmt.Println("golang~")
   case "java":
     fmt.Println("java")
   default:
        fmt.Println("일치하는 값 없음")
 }

// 예제5
switch i,j := 20, 30 ; { //세미콜론 찍어줘야함

case i < j:
  fmt.Println("i는 j보다 작다")
case i == j:
  fmt.Println("i는 j와 같습니다")
case i>j:
  fmt.Println("i는 j보다 크다")
 }//스위치문으로 대소비교, string비교 등등 다양하게 사용이 가능함. 오픈소스에서 스위치문 사용이 빈번함.

}
