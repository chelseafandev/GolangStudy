package explicitcancellation

import (
	"sync"
)

func Gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	return out
}

func Sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// The value type of done is the empty struct because the value doesn’t matter: it is the receive event that indicates the send on out should be abandoned.
func Merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// sync.WaitGroup을 사용하는 이유는?
	// 이미 닫힌 채널에 send 요청을 하는 것은 panic을 발생시키기 때문에, 채널을 닫는 close 함수를 호출하기 전에 모든 send 요청이 끝난다는 것이 보장되어야 한다.
	// sync.WaitGroup 타입은 이러한 동기화 문제를 해결하는 심플한 해결책을 제공한다.
	var wg sync.WaitGroup // sync.WaitGroup는 모둔 Go루틴이 종료될 때까지 대기해야할 때 사용함
	out := make(chan int)

	// lamda와 같은 형태로 함수(output) 정의
	output := func(c <-chan int) {
		defer wg.Done() // 대기 중인 Go루틴의 수행이 종료되는 것을 알려줌
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs)) // WaitGroup에 대기 중인 Go루틴 개수 추가
	for _, c := range cs {
		go output(c)
	}

	// wg.Wait 함수를 실행할때도 go routine을 사용하고 있음
	go func() {
		wg.Wait() // 모든 Go루틴이 종료될 때까지 대기
		close(out)
	}()

	return out
}
