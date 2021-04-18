package main

import (
	"GolangStudy/tuckerlecture/datastruct"
	"fmt"
)

func main() {
	m := datastruct.NewMap()
	m.Add("AAA", "01012345678")
	m.Add("BBB", "01022223333")
	m.Add("CCC", "01033334444")
	m.Add("DDD", "01044445555")

	fmt.Println("AAA =", m.Get("AAA"))
	fmt.Println("BBB =", m.Get("BBB"))
	fmt.Println("CCC =", m.Get("CCC"))
	fmt.Println("FFF =", m.Get("FFF"))
}
