package main

import "fmt"

func recursiveSumV1(num int) int {
	if num == 1 {
		return 1
	}

	return num + recursiveSumV1(num-1)
}

func recursiveSumV2(num int, sum int) int {
	if num == 0 {
		return sum
	} else {
		sum += num
	}

	return recursiveSumV2(num-1, sum)
}

func main() {
	fmt.Println("recursive Version1. 1부터 10까지의 합:", recursiveSumV1(10))
	fmt.Println("recursive Version2. 1부터 10까지의 합:", recursiveSumV2(10, 0))
}
