package main

import (
    "fmt"
    "math"
    myShape "GolangStudy/UnitTest/interface/myshape"
)

type square struct {
    X float64
}

type circle struct {
    R float64
}

func (s square) Area() float64 {
    return s.X * s.X
}

func (s square) Perimeter() float64 {
    return 4 * s.X
}

func (c circle) Area() float64 {
    return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
    return 2 * c.R * math.Pi
}

func Calculate(x myShape.Shape) {
    // 타입 어써션을 통해 인터페이스의 특정 타입을 확인할 수 있다
    _, ok := x.(circle)
    if ok {
        fmt.Println("Is a circle!")
    }

    // 타입 어써션을 통해 인터페이스에 저장된 구체적인 값(여기서는 변수 v에 저장)을 가져올 수 있다
    v, ok := x.(square)
    if ok {
        fmt.Println("Is a square:", v)
    }

    fmt.Println(x.Area())
    fmt.Println(x.Perimeter())
}

func main() {
    x := square{X: 10}
    fmt.Println("Perimeter:", x.Perimeter())
    Calculate(x)

    y := circle{R: 5}
    Calculate(y)
}