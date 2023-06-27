package main

import (
	"fmt"
)

/*
	숫자 : uint8,16,32,64   //unsigned int (양수)
	int8,16,32,64 //정수
	float32,64 //실수

	byte == uint8
	rune == 문자 1개
	int == 플랫폼에 따라 달라짐, 64bit면 int64
	uint = 위와 동일

	그외:
	bool
	string
	배열
	슬라이스
	구조체
	포인터
	함수타입
	맵
	인터페이스
	채널

*/

var g string = "global" //패키지 전역변수

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK    = true
)

func main() {

	var a = 10 //타입추론
	var msg string = "hello Variable"
	c := 5 //var 생략가능, c declared and not used 무조건 선언하면 써야 되네.. 되게 opnioned 하다는 게 이런 거에서 나타나는구나..

	// https://stackoverflow.com/questions/21743841/how-to-avoid-annoying-error-declared-and-not-used
	a = 20
	msg = "good morning"

	// c= 1.1  //타입캐스팅은 어떻게 하나?
	var d float64 = float64(c)
	//캐스팅 시 자료유실 주의

	var f int //기본값 0
	var h string

	if h == "" {
		fmt.Println("h is ''")
	}

	fmt.Println(msg, a, c, d, f, h)

	//PRODUCT = "asd"

	{
		fmt.Println(a)
		var a = 11
		fmt.Println(a)
	}

}
