//반복문 예제
//Go에서는 반복문이 For문만 존재
package main
import "fmt"
func main(){

    //다양한 사용법을 숙지하자
    //예제1
    for i := 0; i < 5; i++ {
      fmt.Println("ex1 : ", i)
    }
//에러발생1
/*
 for i := 0; i < 5; i++   //괄호 내리지말고 첫선언 붙여서~
 {
  }
*/

/* 에러발생 2
for i := 0; i<5; i++    //괄호생략불가
  fmt.Println("ex1:",i)
*/


//예제 2번 (무한루프)
/* for {
  fmt.Print("ex2: 무한루프")
}
// 조건이 없는 for문이 그냥 무한루프문임
*/

 //예제3번 (Range 용법)
loc := []string{"Seoul","Busan","Incheon"}
for index, name := range loc {  //첫번째는 인덱스, 두번째는 값을 가져옴
  fmt.Println("EX3:", index, name)
  }

  eff := []string{"Seoul","Busan","Incheon"}
  for name, index := range eff {  //
    fmt.Println("EX3:", name, index) // name이지만 첫번째 선언된변수라 index를 가져옴
    }
}
