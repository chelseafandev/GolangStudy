/*
Golang 공식 블로그

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines#TOC_3.


Here are the guidelines for pipeline construction:
- stages close their outbound channels when all the send operations are done.
- stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.
*/
package main

import (
	explicitcancellation "GolangStudy/ConcurrencyInGo/pipelines/explicitcancellation"
	fanoutfanin "GolangStudy/ConcurrencyInGo/pipelines/fanoutfanin"
	squaringnumbers "GolangStudy/ConcurrencyInGo/pipelines/squaringnumbers"
	"fmt"
)

func fibonacci(c chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonaccitest() {
	c := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		close(quit)
	}()
	fibonacci(c, quit)
}

func testcase1() {
	// gen -> sq -> main

	// Set up the pipeline.
	c := squaringnumbers.Gen(2, 3)
	out := squaringnumbers.Sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func testcase2() {
	// fan-out
	// 동일한 채널(in)로부터 전달되는 데이터를 2개 이상의 함수(sq)에서 전달받아 처리하는 형태
	in := squaringnumbers.Gen(4, 9)
	c1 := squaringnumbers.Sq(in)
	c2 := squaringnumbers.Sq(in)

	// Consume the merged output from c1 and c2.
	for n := range fanoutfanin.Merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

// explicit cancellation
func testcase3() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := explicitcancellation.Gen(2, 3)

	c1 := explicitcancellation.Sq(done, in)
	c2 := explicitcancellation.Sq(done, in)

	out := explicitcancellation.Merge(done, c1, c2)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func testcase4() {
	fibonaccitest()
}

func main() {
	fmt.Println("== testcase1. squaring numbers ==")
	testcase1()
	fmt.Printf("=================================\n\n")

	fmt.Println("== testcase2. fan-out, fan-in ==")
	testcase2()
	fmt.Printf("=================================\n\n")

	fmt.Println("== testcase3. explicit cancellation ==")
	testcase3()
	fmt.Printf("=======================================\n\n")

	fmt.Println("== testcase4. fibonacci ==")
	testcase4()
	fmt.Printf("=======================================\n\n")

}
