package main
import "fmt"
func main() {
 // 상수
 // const사용을 초기화, 한번 선언 후 값 변경 금지, 고정된 값 관리 용이
 const a string = "Test1"
 const b = "Test2"
 const c int32 = 10 * 10
 // const d = getHeight() 이런식 선언하면 안됨.
 // 왜냐하면 const는 항상 고정된 값이어야하는데, getHeight값의 결과값은 변수일 확률이 있음.
 const e = 35.6
 const f = false
 /*
  에러발생
  const g string
  g = "Test3"        //상수는 선언과 동시에 값이 들어가야함
 */
  fmt.Println("a:",a, "b:",b, "c:",c,  "e:",e, "f:",f)
}
