package goroutine

import (
	"fmt"
	"time"
)

func Greet(msg string, doneChan chan bool) {
	fmt.Println(msg)
	doneChan <- true
}

func SlowGreet(msg string, doneChan chan bool) {
	fmt.Println(msg)
	time.Sleep(3 * time.Second)
	doneChan <- true
}
