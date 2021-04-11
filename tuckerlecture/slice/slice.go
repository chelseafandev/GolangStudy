package main

import (
	"fmt"
)

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2
	fmt.Println("len(a):", len(a))
	fmt.Println("cap(a):", cap(a))

	// append의 입력에 들어가는 slice의 주소값과 리턴값의 주소값은 다르다. 즉, 서로 다른 메모리 공간에 존재한다는 의미!
	b := append(a, 3)
	fmt.Printf("address(a):%p\n", &a)
	fmt.Printf("address(b):%p\n", &b)

	// 같은 원소를 갖고있으나 서로 다른 공간을 확보하고 싶을때는 아래와 같이 처리하자
	c := make([]int, 2, 4)
	c[0] = 1
	c[1] = 2

	d := make([]int, len(c))

	for i := 0; i < len(c); i++ {
		d[i] = c[i]
	}

	d = append(d, 3)
	fmt.Printf("address(c):%p\n", &c)
	fmt.Printf("address(d):%p\n", &d)

	fmt.Println("================================")
	sliceTest()

	fmt.Println("================================")
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var target int
		input, target = RemoveBack(input)
		fmt.Println(target)
	}
	fmt.Println(input)

	input2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var target int
		input2, target = RemoveFront(input2)
		fmt.Println(target)
	}
	fmt.Println(input2)
}

func sliceTest() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	b[0] = 50
	b[1] = 60
	c := a[4:]
	d := a[:4]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// 실제 메모리 상에는 값이 남아있지만 가리키는 포인터만 바뀐다
func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

// 실제 메모리 상에는 값이 남아있지만 가리키는 포인터만 바뀐다
func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
