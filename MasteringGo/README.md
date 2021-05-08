<div align=center>

![](resources/images/masteringgo.jpg)

</div>

## Mastering GO
Mastering GO 정리 노트이며, 챕터 별 핵심 내용을 나름대로 정리해보았다. <br>
다른 언어와 유사한 기초적인 문법에 대한 설명은 제외했으며, 최대한 Go언어만이 갖고 있는 특성에 집중했다. <br>
(챕터 별로 폴더를 생성했고, 해당 폴더 내에 README.md 파일이 포함되어있다. Link를 클릭하면 해당 챕터로 이동한다.)

## 1장 Go 언어와 운영 체제 [Link](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo/chapter1)
- godoc 유틸리티
- 중괄호 작성 스타일을 따를 것
- Go 패키지 다운로드하기
	- Module을 사용해서 빌드하는 연습을 하자😎
	- 외부 패키지 사용하기
- 유닉스 stdin, stdout, stderr
- 커맨드라인 인수 다루기
- Go 언어에서 에러 처리하기
	- 에러 처리하기

## 2장 Go 언어의 내부 살펴보기 [Link](https://github.com/junhaeng90/GolangStudy/tree/main/MasteringGo/chapter2)
- 가비지 컬렉션
- 삼색 알고리즘(tricolor mark-and-sweep alogorithm)
- Go 언어 가비지 컬렉터의 구체적인 작동 방식
  - 삼색 알고리즘의 불변 속성을 유지하기 위한 constraints
- 언세이프(Unsafe) 코드 및 패키지
- defer 키워드
- panic 함수와 recover 함수
- Go 언어의 슬라이스
  - sort.Slice()로 슬라이스 정렬하기
- Go 언어의 상수 생성자 iota