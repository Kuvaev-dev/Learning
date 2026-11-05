package channels

import "fmt"

func getString(c chan string) {
	str := "Hello"
	c <- str
}

func UseChannel() {
	channel := make(chan string)
	go getString(channel)
	hello := <-channel
	fmt.Println(hello)
}

func UseBufferedChannel() {
	channel := make(chan string, 2)
	channel <- "Hello"
	channel <- "World"
	// channel <- "Hey" - Deadlock
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	channel <- "QQ"
	fmt.Println(<-channel)
}
