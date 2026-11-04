package gorutine

import (
	"fmt"
	"time"
)

func doSmth(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func GetRoutine() {
	go doSmth("Hop")
	doSmth("Hey")
}
