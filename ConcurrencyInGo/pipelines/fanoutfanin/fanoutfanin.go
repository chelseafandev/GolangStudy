/*
Fan-out
Multiple functions can read from the same channel until that channel is closed; this is called fan-out.
This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

┏━━━━━━━━━┓
┃ worker1 ┃<────────────┐
┗━━━━━━━━━┛			┌───────┐
					│channel│
┏━━━━━━━━━┓			└───────┘
┃ worker2 ┃<────────────┘
┗━━━━━━━━━┛

Fan-in
A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.

					┌────────┐
					│channel1│
					└───┬────┘
┏━━━━━━━━━━┓			│
┃  worker  ┃<───────────┤
┗━━━━━━━━━━┛			│
					┌───┴────┐
					│channel2│
					└────────┘
*/

package fanoutfanin

import "sync"

// The merge function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies the values to the sole outbound channel.
func Merge(cs ...<-chan int) <-chan int {
	// sync.WaitGroup을 사용하는 이유는?
	// 이미 닫힌 채널에 send 요청을 하는 것은 panic을 발생시키기 때문에, 채널을 닫는 close 함수를 호출하기 전에 모든 send 요청이 끝난다는 것이 보장되어야 한다.
	// sync.WaitGroup 타입은 이러한 동기화 문제를 해결하는 심플한 해결책을 제공한다.
	var wg sync.WaitGroup // sync.WaitGroup는 모둔 Go루틴이 종료될 때까지 대기해야할 때 사용함
	out := make(chan int)

	// lamda와 같은 형태로 함수(output) 정의
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done() // 대기 중인 Go루틴의 수행이 종료되는 것을 알려줌
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
