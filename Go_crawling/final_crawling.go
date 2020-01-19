//최종 실습 예제
//대상 사이트 : 루리웹 (ruliweb.com)
package main
import (
 "bufio"
 "fmt"
  "net/http"
   "os"
  "strings" //parsing
  "sync"
  "github.com/yhat/scrape"
  "golang.org/x/net/html"
  "golang.org/x/net/html/atom"
)
//스크랩핑 대상 URL
const (
  urlRoot = "http://ruliweb.com"
)
//첫 방문(메인페이지) 대상으로 원하는 url 파싱 후 반환하는 함수
//html내 모든 노드 방문하면서 bool return
func parseMainNodes(n *html.Node) bool {
  if n.DataAtom == atom.A && n.Parent != nil {  //a태그 찾고, 부모가 nil 아닐 때
			return scrape.Attr(n.Parent, "class") == "row" //클래스의 row면 긁어와라
		}
    return false //못찾으면 false
}
//에러체크용 공통 함수
func errCheck(err error) {
  if err != nil {
    panic(err)
  }
}
//동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup //wg - working group

//Url 서브페이지 대상으로 원하는내용 parsing 후 반환
func scrapContents(url string, fn string) { //서브게시판은 a태그, class는 deco로 되어있음
    defer wg.Done() //작업종료 알림
//get 방식 요청
  resp, err := http.Get(url)
  errCheck(err)

  //요청 Body 닫기
  defer resp.Body.Close()

  //응답 데이터(html)
   root, err := html.Parse(resp.Body)
   errCheck(err)

   //Response 데이터 html의 원하는 부분 파싱, (서브게시판 부분임)
   matchNode := func(n *html.Node) bool {
     return n.DataAtom == atom.A && scrape.Attr(n, "class") =="deco"
     //페이지 전체 노드를 탐색하며, a태이면서 그 및 클래스가 deco인것만 true로 리턴
   }

  //파일 스트림 생성 -> 파일명, 옵션, 권한
  file, err := os.OpenFile("c:/go_study/scrape/"+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
  //파일 읽기쓰기 권한 부여
  errCheck(err)

  //method 종료시 파일 닫기
  defer file.Close()
  //쓰기 버퍼 선언
  w := bufio.NewWriter(file)

  //matchnode 메소드를 사용해 원하는 노드 순회(Iterator)하면서 출력
  for _, g:= range scrape.FindAll(root, matchNode){
    //Url 및 해당 데이터 출력
    fmt.Println("result : ", scrape.Text(g)) // class가 deco인 text가 출력/
    //파싱 데이터 -> 버퍼에 기록
   w.WriteString(scrape.Text(g)+"\r\n")
  }
 w.Flush() //버퍼 비우면서 기록됨.

}


func main() {
//메인 페이지 get 방식 요청
    resp, err := http.Get(urlRoot) //response,request //응답/요청
    errCheck(err)

    //요청 body 닫기
    defer resp.Body.Close()

    //응답 데이터 (Html)로 날라옴
    root, err := html.Parse(resp.Body) //Body내용 parsing
    errCheck(err)
    //url의 소스코드 내용을 긁어와 Body에 담고, root에 대입.

    //ParseMainNodes 메소드를 크롤링(스크랩핑) 대상 URL 추출
    urlList := scrape.FindAll(root, parseMainNodes) //a태그찾아 파싱한후 urlList에 넘김

    for _, link :=  range urlList {
      //대상 Url 1차 출력
      //fmt.Println("check main link : ", link, idx)// urlList로부터 link, idx값 추출
      //fmt.Println("TargetUrl : ", scrape.Attr(link,"href")) //href URL값만 스크랩하는경우
      fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
      //fmt.Println("fileName:",fileName)
      //해석 -> scrape의 link속성의 href에서, http~ 부분에 해당하는 내용을, ""(공백)으로 replace해라. 1은 바꿀 갯수

      //string.Replace 문법. 예제
       //fmt.Println(strings.Replace("oink oink oink oink","k","ky",2))
       //아웃풋은 oinky oinky oink oink가 나옴.

       //작업 대기열에 추가
       wg.Add(1) //Done 개수와 일치해야 프로그램 종료
       //고루틴 시작 -> 작업 대기열 개수와 같아야한다.
       go scrapContents(scrape.Attr(link, "href"), fileName)
}

//모든 작업이 종료될 때 까지 대기 
       wg.Wait()


    }
