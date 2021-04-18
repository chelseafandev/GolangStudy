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

	fmt.Println()

	// go 내장 map 생성하기: make(map[keytype]valuetype)
	builtinMap := make(map[string]string)
	builtinMap["abc"] = "jhlee"
	fmt.Println(builtinMap["abc"])
	builtinMap["123"] = "jhmoon"
	builtinMap["ASD"] = "john"
	builtinMap["mrt"] = "lee"

	builtinMap2 := make(map[int]int)
	builtinMap2[5] = 0 // 실제로 디폴트 값으로 셋팅하는 경우를 어떻게 구분할까?
	var isExist bool
	var v1, v2 int
	// bool값으로 이 값이 셋팅된건지 아니면 디폴트값인지를 구분할 수 있음
	v1, isExist = builtinMap2[5]
	fmt.Println(v1, isExist)
	v2, isExist = builtinMap2[4]
	fmt.Println(v2, isExist)
	delete(builtinMap2, 5) // map 삭제
	fmt.Println(v1, isExist)

	// map 순회 (key값은 정렬되지 않고 무작위로 나온다!)
	for key, value := range builtinMap {
		fmt.Println("key:", key, "value:", value)
	}
}
