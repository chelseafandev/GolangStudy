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
}
