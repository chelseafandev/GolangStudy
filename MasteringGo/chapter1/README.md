## 1장 Go 언어와 운영 체제

---
- [1장 Go 언어와 운영 체제](#1장-go-언어와-운영-체제)
	- [godoc 유틸리티](#godoc-유틸리티)
	- [중괄호 작성 스타일을 따를 것](#중괄호-작성-스타일을-따를-것)
	- [Go 패키지 다운로드하기](#go-패키지-다운로드하기)
		- [Module을 사용해서 빌드하는 연습을 하자😎](#module을-사용해서-빌드하는-연습을-하자)
		- [외부 패키지 사용하기](#외부-패키지-사용하기)
	- [유닉스 stdin, stdout, stderr](#유닉스-stdin-stdout-stderr)
	- [커맨드라인 인수 다루기](#커맨드라인-인수-다루기)
	- [Go 언어에서 에러 처리하기](#go-언어에서-에러-처리하기)
		- [에러 처리하기](#에러-처리하기)

[뒤로](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo)
<br>

### godoc 유틸리티
godoc을 사용하면 인터넷에 연결돼 있지 않아도 Go 함수나 패키지에 대한 문서를 쉽게 볼 수 있다. <br>
아래 명령어를 통해 별도로 설치해야 사용이 가능하다. <br>
```
go get -v  golang.org/x/tools/cmd/godoc
```
<br>

### 중괄호 작성 스타일을 따를 것
Go 언어는 세미콜론으로 문장의 끝을 표현하며, 컴파일러는 코드에서 필요하다고 판단되는 지점에 <br>
세미콜론을 집어넣는다. 따라서 여는 중괄호로 문장을 시작하면 Go 컴파일러는 그 이전 문장의 끝인 <br>
func main()에 세미콜론을 넣기 때문에 위와 같은 에러 메시지가 발생하는 것이다.
```go
package main

import (
    "fmt"
)
// 에러 발생!
func main()
{
    fmt.Println("Hello World")
}
```

```go
package main

import (
    "fmt"
)
// 정상 동작!
func main() {
    fmt.Println("Hello World")
}
```
<br>

### Go 패키지 다운로드하기
이 책에서 해당 챕터를 집필하는 시점의 stable 버전은 1.9.1 이므로 go module([1.11에 도입](https://golang.org/doc/go1.11#modules))에 대한 설명이 없다. <br>
go module을 기반으로한 로컬 및 외부 패키지 사용 방법을 기술하겠다.

#### Module을 사용해서 빌드하는 연습을 하자😎
로컬에서 작업한 패키지 정보를 사용하는 방법부터 설명하겠다.

1. 아래 명령어를 통해 go module을 항상 사용하겠다고 설정하자.
    ```
    go env -w GO111MODULE=on
    ```
2. go.mod 파일을 추가하자.
    * 추가할 로컬 패키지의 경로에서 아래 명령어를 실행하자.
        ```
        /home/jhlee/go/src/learngo/banking> go mod init
        ```
    * main 함수가 존재하는 경로에서 아래 명령어를 실행하자.
        ```
        /home/jhlee/go/src/learngo> go mod init
        ```
3. go.mod 파일을 수정하자.

    local에 존재하는 패키지를 import 하기 위해서는 <U>replace를 통해 실제 로컬 경로를 추가</U>해주어야한다.
    ```go
    module learngo

    go 1.16

    require learngo/banking v1.16.0

    replace learngo/banking v1.16.0 => ./ banking
    ```

#### 외부 패키지 사용하기
go module을 활용하여 외부 패키지를 사용하는 방법에 대해 설명하겠다.

1. go.mod 수정
    * 추가할 패키지 정보를 go.mod에 추가한다. 여기서는 [gocolly/colly](https://github.com/gocolly/colly)를 예로 들겠다.
    ```go
    module learngo

    go 1.16

    require (
	    github.com/gocolly/colly/v2 v2.1.0
	    learngo/banking v1.16.0
    )

    replace learngo/banking v1.16.0 => ./ banking
    ```

2. 소스에 패키지 정보를 import
    ```go
    package main
    
    import (
        "github.com/gocolly/colly/v2"
    )

    func main() {
	    // Instantiate default collector
	    c := colly.NewCollector(
		    // Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		    colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	    )

	    // On every a element which has href attribute call callback
	    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		    link := e.Attr("href")
		    // Print link
		    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		    // Visit link found on page
		    // Only those links are visited which are in AllowedDomains
		    c.Visit(e.Request.AbsoluteURL(link))
	    })

	    // Before making a request print "Visiting ..."
	    c.OnRequest(func(r *colly.Request) {
		    fmt.Println("Visiting", r.URL.String())
	    })

	    // Start scraping on https://hackerspaces.org
	    c.Visit("https://hackerspaces.org/")
    } 
    ```
3. go mod tidy를 통해 go.sum 파일을 자동으로 채워주자.
    * main 함수가 존재하는 경로에서 아래 명령어를 실행하자.
    ```
    /home/jhlee/go/src/learngo> go mod tidy
    ```
<br>

### 유닉스 stdin, stdout, stderr
기본적으로 모든 유닉스 시스템은 /dev/stdin, /dev/stdout, /dev/stderr라는 세 가지의 특수한 표준 파일명을 사용한다. <br>
각각에 대한 파일 디스크립터는 0, 1, 2다. 이러한 세 가지 파일 디스크립터를 표준 입력, 표준 출력, 표준 에러라고도 부른다. <br>
아래 예제는 표준 에러 출력을 /dev/null 디바이스로 redirection하고 있다. 이는 표준 에러를 무시하겠다는 의미다.
```
go run stdERR.go 2>/dev/null
```
표준 출력(1)과 표준 에러(2)의 결과를 모두 같은 파일에 저장하고 싶다면, 표준 에러에 대한 파일 디스크립터(2)를 표준 출력에 대한 <br>
파일 디스크립터(1)로 redirection하면 된다. 예를 들면 아래와 같다.
```
go run stdERR.go >/tmp/output 2>&1
```
마지막으로 표준 출력과 표준 에러를 모두 /dev/null 디바이스로 보내버릴 수도 있다.
```
go run stdERR.go >/dev/null 2>&1
```
<br>

### 커맨드라인 인수 다루기
커맨드라인 인수를 가져오려면 os 패키지를 이용해야한다. 소스 라인의 주석을 참조바란다.
```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args는 string값을 가진 Go 슬라이스다.
	// 이 슬라이스의 첫 번째 원소(arguments[0])는 실행한 프로그램의 이름이다.
	fmt.Println("len(os.Agrs) =", len(os.Args))
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	fmt.Println("arguments =", arguments)

	min, _ := strconv.ParseInt(arguments[1], 0, 64)
	max, _ := strconv.ParseInt(arguments[1], 0, 64)

	// 입력된 커맨드라인 인수들 중(첫 번째 원소인 프로그램 이름은 제외)에서 최소값과 최대값을 구한다.
	for i := 1; i < len(os.Args); i++ {
		n, _ := strconv.ParseInt(arguments[i], 0, 64)
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
```
<br>

### Go 언어에서 에러 처리하기
새로운 error 변수를 생성하려면 표준 Go 패키지인 errors에서 제공하는 New() 함수를 호출해야 한다.
```go
package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		// 새로운 error 변수를 생성하는 방법
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	// Error() 메소드를 통해 error 변수를 string 타입으로 변환할 수 있음
	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
```

#### 에러 처리하기
앞서 커맨드라인 인수 다루기에서 사용했던 예제 코드를 에러 처리가 가능하도록 개선해보겠다.
```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args는 string값을 가진 Go 슬라이스다.
	// 이 슬라이스의 첫 번째 원소(arguments[0])는 실행한 프로그램의 이름이다.
	fmt.Println("len(os.Agrs) =", len(os.Args))
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more integer.")
		os.Exit(1)
	}

	arguments := os.Args
	// errors.New()를 통해 새로운 error 변수를 생성함
	var err error = errors.New("An error")
	k := 1
	var n int64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a integer!")
			return
		}
		n, err = strconv.ParseInt(arguments[k], 0, 64)
		k++
	}

	min, max := n, n

	// 입력된 커맨드라인 인수들 중(첫 번째 원소인 프로그램 이름은 제외)에서 최소값과 최대값을 구한다.
	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseInt(arguments[i], 0, 64)
		if err == nil {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
```