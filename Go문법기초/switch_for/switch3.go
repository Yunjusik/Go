//스위치문 3
package main

import (
   "fmt"

  )//import 한번에 하기

func main() {
  //예제1, 케이스가 여러경우에 걸리는 예제
  a := 30/15
  switch a {
  case 2,4,6 ://a가 2,4,6 일 경우
  fmt.Println("a->",a,"는 짝수")
  case 1,3,5 : //a는 1,3,5,일 경우

    fmt.Println("a->", a, "는 홀수")
   }

      //fallthrough.

      switch e := "go"; e{
      case "java":
        fmt.Println("java")
        fallthrough
      case "go":
        fmt.Println("go")
        fallthrough
      case "python":
        fmt.Println("python")

      case "ruby":
        fmt.Println("ruby")
        fallthrough
      case "c":
        fmt.Println("c")
// fallthrough 밑에 반드시 실행시켜야하는 문장을 넣어놈
//반드시 의존관계가 있는 명령어를 실행할때 사용
//ex, 요청에 대한 응답 예외처리용
      }
}
