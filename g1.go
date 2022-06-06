package main

import "fmt"

type Student struct {
}

func main() {

	numb := 10
	fmt.Printf("你好世界%d", numb)

	for i := 0; i < 200; i++ {
		fmt.Printf("你好世界%d", i)
	}

}
