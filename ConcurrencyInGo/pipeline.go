/*
Golang 공식 블로그

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines#TOC_3.
*/

package main

import "fmt"

// 첫번째 stage(source)
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

// 두번째 stage
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

// 마지막 stage(sink)
// it receives values from the second stage and prints each one, until the channel is closed

func main() {

	// gen -> sq -> main

	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}
