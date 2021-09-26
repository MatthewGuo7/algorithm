package main

import "fmt"

func main() {
	//ch1 := make(chan int, 10)
	var ch1 chan int
	ch2 := make(chan int)

	ch1 <- 10
	ch2 <- 10

	value, ok := <-ch1
	fmt.Println(value, ok)

	value1 := <- ch2
	fmt.Println(value1)

	select {
	case <- ch1:
	case <- ch2:

	}

	close(ch2)

	select {}
}
