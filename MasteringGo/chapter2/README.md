## 2장 Go 언어의 내부 살펴보기

- [2장 Go 언어의 내부 살펴보기](#2장-go-언어의-내부-살펴보기)
	- [가비지 컬렉션](#가비지-컬렉션)
	- [삼색 알고리즘(tricolor mark-and-sweep alogorithm)](#삼색-알고리즘tricolor-mark-and-sweep-alogorithm)
	- [Go 언어 가비지 컬렉터의 구체적인 작동 방식](#go-언어-가비지-컬렉터의-구체적인-작동-방식)
		- [삼색 알고리즘의 불변 속성을 유지하기 위한 constraints](#삼색-알고리즘의-불변-속성을-유지하기-위한-constraints)
	- [언세이프(Unsafe) 코드 및 패키지](#언세이프unsafe-코드-및-패키지)
	- [defer 키워드](#defer-키워드)
	- [panic 함수와 recover 함수](#panic-함수와-recover-함수)
	- [Go 언어의 슬라이스](#go-언어의-슬라이스)
		- [sort.Slice()로 슬라이스 정렬하기](#sortslice로-슬라이스-정렬하기)
	- [Go 언어의 상수 생성자 iota](#go-언어의-상수-생성자-iota)

[뒤로](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo)
<br>

### 가비지 컬렉션
가비지 컬렉션이란 더 이상 사용하지 않는 메모리 공산을 해제하는 프로세스다. 다시 말해 가비지 컬렉터는 <br>
현재 스코프를 벗어난 오브젝트를 발견하고, 이를 더 이상 참조할 일이 없다고 판단하면 그 공간을 해제한다.
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	// 가비지 컬렉션 통계에 대한 최신 정보를 조회하려면, 매번 runtime.ReadMemStats() 함수르 호출해야한다.
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}
```

Go 가비지 컬렉터의 작동 과정을 보다 상세하게 출력하려면 아래와 같이 실행한다. 이를 통해 가비지 컬렉션 프로세스가 <br>
진행되는 동안 할당된 힙 크기를 보다 자세히 알 수 있다.
```
GODEBUG=gctrace=1 go run main.go
```

예를 들어 190->190->0 MB라고 표현된 부분을 보자. 여기서 첫 번째 숫자는 가비지 컬렉터가 실행될 시점의 힙 크기다. <br>
두 번째 값은 가비지 컬렉터가 실행을 마칠 시점의 힙 크기다. 마지막은 현재 힙 크기다. <br>

gc 17 @30.304s 0%: 0.015+0.15+0.002 ms clock, 0.12+0/0.13/0+0.016 ms cpu, <span style="color:red">190->190->0 MB</span>, 191 MB goal, 8 P <br>
gc 18 @35.338s 0%: 0.026+0.34+0.003 ms clock, 0.21+0/0.28/0.063+0.030 ms cpu, <span style="color:red">95->95->0 MB</span>, 96 MB goal, 8 P <br>
gc 19 @40.360s 0%: 0.015+0.14+0.003 ms clock, 0.12+0/0.12/0.002+0.024 ms cpu, <span style="color:red">95->95->0 MB</span>, 96 MB goal, 8 P <br>
<br>

### 삼색 알고리즘(tricolor mark-and-sweep alogorithm)
Go 언어의 가비지 컬렉터는 삼색 알고리즘에 따라 작동한다. 삼색 알고리즘의 핵심 원리는 힙에 있는 오브젝트를 <br>
이 알고리즘에 따라 세 가지 색깔로 지정된 집합으로 나누는데 있다.
* 흰색 집합(white set): 프로그램에서 더 이상 접근할 수 없어서 가비지 컬렉션 대상이 되는 오브젝트
* 회색 집합(gray set): 프로그램이 현재 사용하고 있지만 흰색 오브젝트를 가리킬 수도 있어서 검사 과정을 거쳐야하는 오브젝트
* 검은색 집합(black set): 프로그램이 사용하고 있으며, 흰색 집합의 오브젝트를 가리키는 포인터가 확실히 없는 오브젝트

여기서 주목할 점은 검은색 집합에서 곧바로 흰색 집합으로 갈 수 없으며, 검은색 집합에 있는 오브젝트는 흰색 집합의 오브젝트를 직접 가리킬 수 없다.
<br>

### Go 언어 가비지 컬렉터의 구체적인 작동 방식
1) 프로그램의 실행을 잠시 멈춘다.
2) 프로그램의 힙에서 접근할 수 있는 오브젝트를 모두 방문한 뒤에 적절히 표시(mark)한다(흰색, 회색, 검은색 중 하나로 표시).
3) 접근할 수 없는 오브젝트를 쓸어(sweep) 담는다.

#### 삼색 알고리즘의 불변 속성을 유지하기 위한 constraints
* 새로운 오브젝트는 반드시 회색 집합으로 가야한다.
* 프로그램에 나온 포인터가 이동하면 포인터가 가리키던 오브젝트도 회색으로 변경된다.
* 포인터가 변경될 때마다 쓰기 장벽(write barrier)이라 부르는 Go 코드가 자동으로 실행되면서 색깔 변경 작업을 수행한다. <br>
  (쓰기 장벽 코드를 실행함으로써 발생하는 지연 시간은 가비지 컬렉터를 동시에 실행하기 위해 치러야할 대가인 셈)
<br>

### 언세이프(Unsafe) 코드 및 패키지
언세이프 코드란 Go 언어의 타입 안정성 및 메모리 보안 검사를 거치지 않는 코드를 말한다. <br>
언세이프 코드라고 하면 대부분 포인터에 관련된 코드를 의미한다.

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := []int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")

	// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
	//
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()

	// 메모리 주소에 존재하지 않는 배열의 원소에 접근하고 있으나
	// 이 코드에서는 unsafe 패키지를 사용하고 있기 때문에 Go 컴파일러에서 이런 논리적 에러를 찾아 주지 않는다.
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}
```
<br>

### defer 키워드
어떤 함수(A)를 호출하는 문장 앞에 defer 키워드를 붙이면, 이 defer문을 담고 있는 함수가 리턴될 때까지 함수(A)의 실행을 미룬다. <br>
defer문을 담고 있는 함수가 리턴된 후에, defer 키워드를 이용해 실행이 미뤄진 함수가 호출되는 순서는 LIFO방식을 따른다.

```go
package main

import (
	"fmt"
)

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d3()
}
```
<br>

### panic 함수와 recover 함수
pnaic() 함수는 Go 언어에서 기본으로 제공하는 것으로, 기존 Go 프로그램의 실행 흐름을 종료하고 패닉 상태에 빠진다. <br>
recover() 함수 역시 Go언에서 기본으로 제공하며 panic()으로 인해 패닉 상태에 빠졌던 Go 루틴으로부터 제어권을 다시 가져오게 해준다.

```go
package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}
```
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