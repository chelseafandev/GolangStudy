package main

import "fmt"

func main() {
	// (배열 복사 없이) 해당 배열을 역순으로 만들고 싶다
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("original array:", arr)

	for i := 0; i < len(arr)/2; i++ {
		// 이게 swap이 되네??!!
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("reversal array:", arr)
}
