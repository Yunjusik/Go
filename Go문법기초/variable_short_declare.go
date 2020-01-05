package main
import "fmt"
func main(){
//중요*** Go에만 있음
//짧은선언 기능 - > 반드시 함수 안에서만 사용. 전역사용 불가능...
//선언 후 할당하면 예외발생(에러)
//짧은선언은 해당 함수에서만 생성되었다가 사라지는 함수임. 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있음.
//짧선 예제
 shortVar1 := 3
 shortVar2 := "Test"
 shortVar3 := false
 //여기서 shortVar1 :=10, 이런식으로 재할당해주면 오류뜸. 짧선은 한번 선언하면 재선언하지말아야합니다.
//shortVar3 := true   //마찬가지임

 fmt.Println("shortVar1:",shortVar1,"shortVar2:",shortVar2,"shortVar3:",shortVar3)
//짧은선언을 통한 if문 선언

if i:=10; i<11 {  //이런 짧선을 보고 i는 해당 루프에서만 사용되는구나 라는걸 캐치
  fmt.Println("short variable test ~~")
}


}
