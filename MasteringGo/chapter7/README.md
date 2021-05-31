## 7장 리플렉션과 인터페이스

- [7장 리플렉션과 인터페이스](#7장-리플렉션과-인터페이스)
  - [타입 메소드](#타입-메소드)
  - [Go 언어의 인터페이스](#go-언어의-인터페이스)

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