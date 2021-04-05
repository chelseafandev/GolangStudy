package main

import "fmt"

func main() {

	// radix sort 조건
	// 1. 배열의 원소의 범위가 제한적이어야 한다
	// 2. 배열의 원소가 정수 값이어야 한다

	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	tmp := [10]int{}
	sortedArr := [11]int{}

	for i := 0; i < len(arr); i++ {
		tmp[arr[i]]++
	}

	idx := 0
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < tmp[i]; j++ {
			sortedArr[idx] = i
			idx++
		}
	}

	fmt.Println(sortedArr)
}
