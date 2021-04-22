## 2장 Go 언어의 내부 살펴보기

- [2장 Go 언어의 내부 살펴보기](#2장-go-언어의-내부-살펴보기)
  - [가비지 컬렉션](#가비지-컬렉션)
  - [삼색 알고리즘](#삼색-알고리즘)

[뒤로](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo)

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

### 삼색 알고리즘