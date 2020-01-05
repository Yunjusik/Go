//변수1
package main
import "fmt"
func main() {
  //첫 시작시 초기화
  //정수타입 : 0, 실수(소수점):0.0, 문자열: "", Boolean: True, False
  // 변수명: 숫자 첫글자 불가, 대소문자 구분함, 문자 숫자 밑줄 특수기호 사용 가능
  // 변수 및 상수: 함수 내외 사용 가능
  var a int // 변수선언, 변수명, 자료형 이런순서대로정의
  var b string
  var c,d,e int
  var f, g, h int = 1,2,3 //선언과 동시에 초기화 가능
  var i float32 = 11.4
  var j string = "Hellow world~~"
  var k = 4.74 //자료형 선언 안했지만, go가 알아서 float형으로 할당해버림
  var l = "hi jusik" //마찬가지로 string으로 알아서 초기화해줌
  var m = true

  a=4
  b="Hello Go!"
  e=77


  //선언만해놓고 초기화를 안한 변수가 있으면 에러가뜸 (a,b,c,d,e 등등)
fmt.Println("a : ", a) //값을 대입하지않은 int는 0으로초기화
fmt.Println("b : ", b) //string은 null
fmt.Println("c : ", c)
fmt.Println("d : ", d)
fmt.Println("e : ", e)
fmt.Println("l : ", l)
fmt.Println("m : ", m)
fmt.Println("f : ", f)
fmt.Println("g : ", g)
fmt.Println("h : ", h)
fmt.Println("i : ", i)
fmt.Println("j : ", j)
fmt.Println("k : ", k)



}
