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

	// waiter.Add(1)
	go func(message string) {
		say(message) // message에서 "안녕 세상!" 받아서 say로 전달 & 출력
		defer waiter.Done()
	}("안녕 세상!")

	// sync.WaitGroup를 사용해서 고루틴 실행이 끝날 때까지 기다리게 한다.
	// 희한한 건 sync.WaitGroup 대신 아래의 그냥 say()를 실행해도 고루틴이 실행 된다.
	waiter.Wait()
	// say("End")
}
