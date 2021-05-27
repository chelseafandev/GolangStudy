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
      - [Binary Tree](#binary-tree)
      - [DFS](#dfs)
      - [BFS](#bfs)
    - [Heap](#heap)
    - [Map](#map)
  - [Thread](#thread)
  - [Channel](#channel)
<br>

## 함수
* Function [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/function/function.go)

    * func 라는 예약어로 시작하고 함수의 파라미터는 (변수명1 타입, 변수명2 타입)와 같이 변수명 뒤에 타입을 적어주고 콤마(,)를 통해 구분한다.
    * 리턴 값이 존재하는 경우에는 파라미터 뒤에 리턴 값의 타입을 적어준다.

* Recursive [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/recursive/recursive.go)

## 배열과 문자열
* Array [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/array/array.go)
    
    * go에서 Array 타입도 제공하긴 하지만 그냥 마음 편하게 Slice를 쓰자😋

* String [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/string/string.go)

    * 한글은 UTF-8(3byte)로 표현된다.<br>
    예를 들어 "Hello 월드"라는 문자열을 변수 s에 저장하게 되면 월이라는 한글은 s[6], s[7], s[8] 3개의 바이트 변수를 차지하는 것을 확인할 수 있다.
    * go에서는 UTF-8값을 저장하는 타입 rune을 제공하고 있다.<br>
    rune 타입의 변수에 문자열을 저장하게 되면 한글이던 영문이던 동일한 사이즈를 차지하는 것을 볼 수 있으며, string 함수의 인자로 rune 타입 변수를 넘기면 string으로 알아서 변환도 해준다.

* Radix Sort [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/radix/radix.go)

    * radix sort를 사용하기 위해서는 아래 2가지 조건이 필요하다.
        > 1. 배열의 원소의 범위가 제한적이어야 한다. (왜냐? 원소의 범위만큼 배열을 잡아 사용해야하기 때문에)
        > 2. 배열의 원소가 정수 값이어야 한다. (왜냐? 배열의 원소 값이 1번에서 잡아서 사용하는 배열의 인덱스로 활용되기 때문)

## 구조체
* Structure [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/structure/structure.go)

    * go에서는 C언어와 유사한 형태로 구조체를 선언하거나 참조하고 있으며, 클래스를 정의할 때도 구조체를 활용한다.

## 포인터
* Pointer [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/pointer/pointer.go)

    * go에서도 포인터 개념(메모리의 주소값)이 있긴있지만 C언어에서처럼 포인터 연산까지 제공하지는 않는다.
    * 구조체를 통해 클래스를 만들고 해당 클래스에 대한 멤버 함수를 정의할 때 포인터 개념이 반드시 필요하다.

## Slice
* Slice [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/slice/slice.go)

## Instance
* Instance [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/instance/instance.go)

## Data Structure
go언어로 다양한 자료구조를 구현 [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct)
### Linked List & Doubly Linked List
* Linked List & Doubly Linked List [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/linkedlist.go)

### Stack
* Stack [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/stack.go)

### Queue
* Queue [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/queue.go)

### Tree
* Tree [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/tree.go)
#### Binary Tree
* Binary Tree [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/binarytree.go)
#### DFS
#### BFS

### Heap
* Heap [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/heap.go)

### Map
* Map [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/datastruct/map.go)

## Thread
* Thread [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/thread/thread.go)

* BankAccount 예제를 통한 lock의 개념 숙지 [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/bankaccount/bankaccount.go)

## Channel
* channel [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/channel/channel.go)

* select 구문 [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/select/select.go)

* CarFactory 예제를 통한 channel의 개념 숙지 [소스링크](https://github.com/junhaeng90/GolangStudy/main/tuckerlecture/carfactory/carfactory.go)