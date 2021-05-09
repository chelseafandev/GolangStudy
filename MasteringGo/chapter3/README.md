## 3장 Go 언어의 기본 데이터 타입

- [3장 Go 언어의 기본 데이터 타입](#3장-go-언어의-기본-데이터-타입)
	- [Go 언어의 슬라이스](#go-언어의-슬라이스)
		- [sort.Slice()로 슬라이스 정렬하기](#sortslice로-슬라이스-정렬하기)
	- [Go 언어의 상수 생성자 iota](#go-언어의-상수-생성자-iota)

[뒤로](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo)
<br>

### Go 언어의 슬라이스
슬라이스를 함수로 전달할 때 참조값(reference)을 전달하여 함수 안에서 슬라이스에 대해 수정한 사항은 그 함수가 종료되더라도 그대로 유지된다. <br>
또한 함수로 전달 시 메모리 주소만 전달하기 때문에 배열로 전달할 때 보다 훨씬 빠르다.

#### sort.Slice()로 슬라이스 정렬하기
``` go
package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 34})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 30})

	fmt.Println("0:", mySlice)

	// 구체적인 정렬 방식은 익명 함수로 정의하고 있음
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	// 구체적인 정렬 방식은 익명 함수로 정의하고 있음
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
```
<br>

### Go 언어의 상수 생성자 iota
상수 생성자인 iota는 서로 관련된 값들이 순차적으로 증가하도록 선언할 때 사용한다.
```go
package main

import "fmt"

type Power2 int

func main() {
	const (
		p2_0 Power2 = 1 << iota // 최초 iota값은 0이고, 1을 왼쪽으로 0번 시프트해서 나온 값인 1을 p2_0값으로 지정
		_ // iota값이 1인 경우에는 건너뜀
		p2_2 // iota값은 2이고, 1을 왼쪽으로 2번 시프트해서 나온 값인 4을 p2_2값으로 지정
		_ // iota값이 3인 경우에는 건너뜀
		p2_4 // iota값은 4이고, 1을 왼쪽으로 4번 시프트해서 나온 값인 16을 p2_2값으로 지정
		_ // iota값이 5인 경우에는 건너뜀
		p2_6 // iota값은 6이고, 1을 왼쪽으로 6번 시프트해서 나온 값인 64을 p2_2값으로 지정
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^0:", p2_2)
	fmt.Println("2^0:", p2_4)
	fmt.Println("2^0:", p2_6)
}
```
<br>

###