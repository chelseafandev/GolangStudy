/*
Golang 공식 블로그

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines#TOC_3.
*/
package main

import (
	fanoutfanin "GolangStudy/ConcurrencyInGo/pipelines/fanoutfanin"
	squaringnumbers "GolangStudy/ConcurrencyInGo/pipelines/squaringnumbers"
	"fmt"
)

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

func main() {
	fmt.Println("== testcase1. squaring numbers ==")
	testcase1()
	fmt.Printf("=================================\n\n")

	fmt.Println("== testcase2. fan-out, fan-in ==")
	testcase2()
	fmt.Printf("=================================\n\n")
}
