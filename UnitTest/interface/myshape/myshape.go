package myShape

type Shape interface {
    Area() float64 // 도형의 면적 계산
    Perimeter() float64 // 도형의 둘레 계산
}