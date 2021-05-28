<div align=center>

![](resources/images/tuckergo.jpg)

</div>

# 컴맹을 위한 프로그래밍 기초 강좌 [Link](https://www.youtube.com/playlist?list=PLy-g2fnSzUTAaDcLW7hpq0e8Jlt7Zfgd6)

- [컴맹을 위한 프로그래밍 기초 강좌 Link](#컴맹을-위한-프로그래밍-기초-강좌-link)
  - [함수](#함수)
  - [배열과 문자열](#배열과-문자열)
  - [구조체](#구조체)
  - [포인터](#포인터)
  - [Slice](#slice)
  - [Instance](#instance)
  - [Data Structure](#data-structure)
    - [Linked List & Doubly Linked List](#linked-list--doubly-linked-list)
    - [Stack](#stack)
    - [Queue](#queue)
    - [Tree](#tree)
      - [DFS](#dfs)
      - [BFS](#bfs)
      - [Binary Tree](#binary-tree)
    - [Heap](#heap)
    - [Map](#map)
  - [Thread](#thread)
  - [Channel](#channel)
  - [Interface](#interface)
<br>

## 함수
* Function [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/function/function.go)

    * func 라는 예약어로 시작하고 함수의 파라미터는 (변수명1 타입, 변수명2 타입)와 같이 변수명 뒤에 타입을 적어주고 콤마(,)를 통해 구분한다.
    * 리턴 값이 존재하는 경우에는 파라미터 뒤에 리턴 값의 타입을 적어준다.

* Recursive [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/recursive/recursive.go)

## 배열과 문자열
* Array [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/array/array.go)
    
    * go에서 Array 타입도 제공하긴 하지만 그냥 마음 편하게 Slice를 쓰자😋

* String [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/string/string.go)

    * 한글은 UTF-8(3byte)로 표현된다.<br>
    예를 들어 "Hello 월드"라는 문자열을 변수 s에 저장하게 되면 월이라는 한글은 s[6], s[7], s[8] 3개의 바이트 변수를 차지하는 것을 확인할 수 있다.
    * go에서는 UTF-8값을 저장하는 타입 rune을 제공하고 있다.<br>
    rune 타입의 변수에 문자열을 저장하게 되면 한글이던 영문이던 동일한 사이즈를 차지하는 것을 볼 수 있으며, string 함수의 인자로 rune 타입 변수를 넘기면 string으로 알아서 변환도 해준다.

* Radix Sort [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/radix/radix.go)

    * radix sort를 사용하기 위해서는 아래 2가지 조건이 필요하다.
        > 1. 배열의 원소의 범위가 제한적이어야 한다. (왜냐? 원소의 범위만큼 배열을 잡아 사용해야하기 때문에)
        > 2. 배열의 원소가 정수 값이어야 한다. (왜냐? 배열의 원소 값이 1번에서 잡아서 사용하는 배열의 인덱스로 활용되기 때문)

## 구조체
* Structure [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/structure/structure.go)

    * go에서는 C언어와 유사한 형태로 구조체를 선언하거나 참조하고 있으며, 클래스를 정의할 때도 구조체를 활용한다.

## 포인터
* Pointer [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/pointer/pointer.go)

    * go에서도 포인터 개념(메모리의 주소값)이 있긴있지만 C언어에서처럼 포인터 연산까지 제공하지는 않는다.
    * 구조체를 통해 클래스를 만들고 해당 클래스에 대한 멤버 함수를 정의할 때 포인터 개념이 반드시 필요하다.

## Slice
* Slice [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/slice/slice.go)

    * Slice는 실제 메모리 공간을 가리키는 포인터다.<br>
    Slice는 포인터 개념(주소값)이다. slice의 append 함수를 잘 살펴보면 append의 입력에 들어가는 slice의 주소값과 리턴값의 주소값이 다르다는 것을 확인할 수 있는데, 이는 입력의 slice와 출력의 slice는 서로 다른 메모리 공간에 존재한다는 의미😁
    * 실제로 데이터가 저장되는 공간은 별도의 메모리 공간에 존재하고 이를 가리키고 있는게 Slice라고 생각하면 된다.

## Instance
* Instance [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/instance/instance.go)

    * OOP 형태(Subject, Verb, Object)<br>
    컴퓨터 입장에서는 Procedure나 OOP나 동일하지만 프로그래머에게는 완전히 새로운 패러다임(!) OOP의 경우에는 Entity간의 Relationship이 중요하다.
    * 인스턴스란?<br>
    ```go
    package main

    import "fmt"

    type Student struct {
	    name  string
	    age   int
	    grade int
    }

    func (t *Student) SetName(name string) {
	    t.name = name
    }

    func SetName(t *Student, name string) {
	    t.name = name
    }

    func main() {
        var a *Student
        a = &Student{name: "junhaeng", age: 32, grade: 100}

        // 프로시저 기반
	    // 컴퓨터 입장에서는 Procedure든 OOP든 차이가 없다(!)
	    // 아래 a.SetName과 완전히 동일함
	    SetName(a, "jihee")

        // 객체 기반
	    a.SetName("helloworld") // Subject역할을 하는 a를 Instance라고 하자(일종의 생명주기)
    }
    ```
## Data Structure
go언어로 다양한 자료구조를 구현 [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct)
### Linked List & Doubly Linked List
* Linked List & Doubly Linked List [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/linkedlist.go)

### Stack
* Stack [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/stack.go)

### Queue
* Queue [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/queue.go)

### Tree
* Tree [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/tree.go)
#### DFS
* 깊이우선탐색(Depth First Search)

    * Stack을 활용(FILO)<br>
    최초 트리의 루트 노드를 Stack에 저장하고, Stack에 값이 없을 때까지 아래 과정을 반복한다.(FILO 성질을 활용)
    > 1. Stack에서 값을 Pop한다.
    > 2. Pop한 노드의 값을 출력(또는 다른 연산)한다.
    > 3. Pop한 노드의 모든 자식 노드를 Stack에 Push한다.
    > 4. Pop한 노드의 자식 노드가 없다면 Push없이 다시 1번으로 돌아간다.

#### BFS
* 너비우선탐색(Breadth First Search)

    * Queue를 활용(FIFO)<br>
    최초 트리의 루트 노드를 Queue에 저장하고, Queue에 값이 없을 때까지 아래 과정을 반복한다.(FIFO 성질을 활용)
    > 1. Queue에서 값을 Pop한다.
    > 2. Pop한 노드의 값을 출력(또는 다른 연산)한다.
    > 3. Pop한 노드의 모든 자식 노드를 Queue에 Push한다.
    > 4. Pop한 노드의 자식 노드가 없다면 Push없이 다시 1번으로 돌아간다.

#### Binary Tree
* Binary Tree [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/binarytree.go)

### Heap
* Heap [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/heap.go)

### Map
* Map [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/datastruct/map.go)

## Thread
* Thread [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/thread/thread.go)

* BankAccount 예제를 통한 lock의 개념 숙지 [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/bankaccount/bankaccount.go)

## Channel
* channel [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/channel/channel.go)

    * goroutine 간의 데이터 송수신을 위해 사용한다.
    * 채널로 보내는 경우 표기법
    ```go
    채널변수 <- 보낼값
    ```
    * 채널로부터 받는 경우 표기법
    ```go
    받을변수 <- 채널변수
    ```
* select 구문 [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/select/select.go)

* CarFactory 예제를 통한 channel의 개념 숙지 [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/carfactory/carfactory.go)

## Interface
* sandwitch 예제를 통한 Interface의 개념 숙지 [소스링크](https://github.com/junhaeng90/GolangStudy/blob/main/tuckerlecture/sandwitch/sandwitch.go)
  
    * go에서 interface를 사용하기 위해서는 이 interface를 사용할 곳에서 interface내에 정의한 함수를 구현하기만 하면된다.