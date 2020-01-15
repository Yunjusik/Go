// 라이브러리 접근제어(1)
package lib

func CheckNumb(c int32) bool { //여기서 CheckNumb를 checkNumb 소문자로 바꿔서하면 private이 되므로, 불러쓸수가없음... 대문자로 하는것 기억
	return c > 10
}
