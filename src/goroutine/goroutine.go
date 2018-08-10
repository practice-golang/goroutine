package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s + ":" + strconv.Itoa(i))
	}
}

func main() {
	var waiter sync.WaitGroup

	say("Begin!!")

	waiter.Add(1)
	message := "Hello world!"
	go func() {
		say(message)
		defer waiter.Done()
	}()

	waiter.Add(1)
	go func(message string) {
		say(message) // message에서 "안녕 세상!" 받아서 say로 전달 & 출력
		defer waiter.Done()
	}("안녕 세상!")

	waiter.Wait()
}
