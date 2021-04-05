package main

import "fmt"

func main() {
	// 한글은 UTF-8(3byte)로 표현
	s := "Hello 월드"

	// "월" 이라는 한글은 아래 3개의 byte로 구성됨
	fmt.Println(s[6], s[7], s[8])

	// "드" 라는 한글은 아래 3개의 byte로 구성됨
	fmt.Println(s[9], s[10], s[11])

	// go에서는 UTF-8값을 저장하는 타입 rune을 제공함
	s2 := []rune(s)
	fmt.Println("rune array size:", len(s2))
	fmt.Println(s2[6], string(s2[6]))
	fmt.Println(s2[7], string(s2[7]))
}
