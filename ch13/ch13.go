package main

import (
	"errors"
	"fmt"
)

/**
error

http://cloudrain21.com/golang-graceful-error-handling
https://dev-yakuza.posstree.com/ko/golang/error-handling/

*/

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func f3(arg int) (int, error) {
	if arg == 42 {
		panic("a problem")
	}
	return arg + 3, nil
}

func recoverTest(arg int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	f3(arg)

}

func main() {

	//for _, i := range []int{7, 42} {
	//	if r, e := f1(i); e != nil {
	//		fmt.Println("f1 failed:", e)
	//	} else {
	//		fmt.Println("f1 worked:", r)
	//	}
	//}
	//
	//for _, i := range []int{7, 42} {
	//	if r, e := f2(i); e != nil {
	//		fmt.Println("f2 failed:", e)
	//	} else {
	//		fmt.Println("f2 worked:", r)
	//	}
	//}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	recoverTest(42)

	fmt.Println("process end")

}
