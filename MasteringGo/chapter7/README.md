## 7장 리플렉션과 인터페이스

- [7장 리플렉션과 인터페이스](#7장-리플렉션과-인터페이스)
  - [타입 메소드](#타입-메소드)
  - [Go 언어의 인터페이스](#go-언어의-인터페이스)
  - [인터페이스 직접 정의하기](#인터페이스-직접-정의하기)
  - [Go 인터페이스 사용하기](#go-인터페이스-사용하기)
  - [Go 언어에서 OOP하기](#go-언어에서-oop하기)

[뒤로](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo)
<br>

### 타입 메소드
타입 메소드란 특수한 수신자(receiver) 인수를 받는 함수다. 일반 함수처럼 선언하되 함수 이름 앞에 특수한 매개변수가 추가된다.<br>
아래 함수에서 f 매개변수가 메소드의 수신자다. 수신자의 작동 과정을 OOP 용어로 표현하면 오브젝트(Object)에 메시지를 보낸다고 말한다.<br>
Go 언어에서 메소드의 수신자는 변수 이름처럼 정의하며, 흔히 한 문자로 표현한다.
```go
func (f *File) Close() error {
    if err := f.checkValid("close"); err != nil {
        return err
    }
    return f.file.close()
}
```
<br>

### Go 언어의 인터페이스
Go 언어에서 interface 타입을 정의할 때 구체적인 동작을 구현할 메소드의 집합을 나열하는 방식으로 표현한다.<br>
어떤 타입이 특정한 인터페이스를 따르기 위해서, 그 인터페이스에서 정의한 모든 메소드를 구현해야 한다.<br><br>
> 💡 함수의 매개변수 타입을 인터페이스로 정의했을 때 그 인터페이스를 구현한 변수라면 어떤 것도 그 함수에 전달할 수 있다.<br>
> 인터페이스는 특별한 기능을 직접 제공하지 않고 단지 형식을 맞추는 역할만 한다.

Go 인터페이스 중에서도 가장 흔히 사용하는 두가지는 io.Reader와 io.Writer
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
<br>

### 인터페이스 직접 정의하기

```go
package myInterface

type Shape interface {
    Area() float64 // 도형의 면적 계산
    Perimeter() float64 // 도형의 둘레 계산
}
```

myInterface 설치 과정은 아래 커맨드를 통해 진행된다.
```
$ mkdir ~/go/src/myInterface
$ cp myInterface.go ~/go/src/myInterface
$ go install myInterface
```
<br>

### Go 인터페이스 사용하기

```go
package main

import (
    "fmt"
    "math"
    "myInterface"
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

func Calculate(x myInterface.Shape) {
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
```
<br>

### Go 언어에서 OOP하기
Go 언어는 상속을 지원하지 않는 대신 합성(composition)을 지원하고, 인터페이스로 다형성(polymophism)을 지원한다. 따라서 Go 언어는 OOP 언어는 아니지만, 어느정도 OOP를 흉내낼 수 있는 기능은 몇 가지 갖추고 있다.

<br>

첫번째 기법은 타입에 함수를 추가해서 메소드를 정의하는 것이다. 다시 말해 이렇게 작성한 함수와 타입은 일종의 오브젝트인 셈이다.
두번째 기법은 새로 정의한 구조체 타입에 다른 타입을 집어 넣어서 일종의 계층을 형성하는 것이다.

```go
package main

import (
    "fmt"
)

type a struct {
    XX int
    YY int
}

type b struct {
    AA string
    XX int
}

type c struct {
    // 합성(composition) 기능을 이용하여 c라는 데이터 타입으로 a 타입 변수와 b 타입 변수를 하나로 묶고 있다
    A a
    B b
}

// a 타입의 변수에 A() 함수를 정의
func (A a) A() {
    fmt.Println("Function A() for A")
}

// b 타입의 변수에 A() 함수를 정의
// 서로 다른 타입에 속한 함수의 이름을 똑같이 표현할 수 있다
func (B b) A() {
    fmt.Println("Function A() for B")
}


func main() {
    var i c
    i.A.A()
    i.B.A()
}
```