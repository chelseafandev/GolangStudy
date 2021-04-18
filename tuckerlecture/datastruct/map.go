package datastruct

import "fmt"

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		fmt.Println("s[i] =", int(s[i]))
		h = (h*A + int(s[i])) % B
	}

	return h
}
