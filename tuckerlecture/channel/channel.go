package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c // 값을 넣어줄때까지 대기
	fmt.Println(v)
}

func main() {
	var c chan int
	c = make(chan int) // 길이가 0개인 경우에는 다른 곳에서 Pop을 할때까지 대기한다
	go pop(c)
	c <- 10
	//v := <-c
	//fmt.Println(v)
	fmt.Println("End of program!")
}
