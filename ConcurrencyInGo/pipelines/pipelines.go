/*
Golang 공식 블로그

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines#TOC_3.


Fan-out
Multiple functions can read from the same channel until that channel is closed; this is called fan-out.
This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

Fan-in
A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.
*/
package main

import (
	"fmt"
	"sync"
)

// a function that converts a list of integers to a channel that emits the integers in the list.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// receives integers from a channel and returns a channel that emits the square of each received integer.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// The merge function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies the values to the sole outbound channel.
func merge(cs ...<-chan int) <-chan int {
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

func squaringnumbers() {
	// gen -> sq -> main

	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func fanoutfanin() {
	// fan-out
	// 동일한 채널(in)로부터 전달되는 데이터를 2개 이상의 함수(sq)에서 전달받아 처리하는 형태
	in := gen(4, 9)
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func main() {
	fmt.Println("== testcase1. squaring numbers ==")
	squaringnumbers()
	fmt.Printf("=================================\n\n")

	fmt.Println("== testcase2. fan-out, fan-in ==")
	fanoutfanin()
	fmt.Printf("=================================\n\n")
}
