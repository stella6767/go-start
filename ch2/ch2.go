package main

import (
	"fmt"
)

/**
표준 입출력
*/

func main() {

	var a int

	n, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a)
	}

	//stdin := bufio.NewReader(os.Stdin)

}
