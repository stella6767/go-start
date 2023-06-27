package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

	wg.Done()
}

/**
고루틴
Golang이 사용하는 경량 쓰레드

 메인 고루틴이 종료되면 메인 고루틴에서 실행한 서브 고루틴도 모두 함께 종료되게 됩니다.

Go 채널은 그 채널을 통하여 데이타를 주고 받는 통로라 볼 수 있는데, 채널은 make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 을 통해 데이타를 보내고 받는다. 채널은 흔히 goroutine들 사이 데이타를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용된다.
http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90
https://dev-yakuza.posstree.com/ko/golang/goroutine/
*/

func main() {

	runtime.GOMAXPROCS(2) //다중 cpu

	wg.Add(3)

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}("going")

	//time.Sleep(time.Second)
	fmt.Println("done")

	messages := make(chan string)
	go func() {
		messages <- "ping"
		wg.Done()
	}()
	msg := <-messages
	fmt.Println(msg)

	wg.Wait()

	fmt.Println("main goroutine end")

}
