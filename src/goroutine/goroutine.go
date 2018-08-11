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

	// 고루틴은 본체는 따로 돌기 때문에 본체는 자기 할 일을 마치면 goroutine의 동작에 관계없이 종료된다고 한다.
	// 여기서는 sync.WaitGroup를 사용해서 고루틴 실행이 끝날 때까지 기다리게 한다.
	waiter.Wait()

	// 본체가 고루틴보다 늦게 끝난다면 sync 사용이 없어도 고루틴의 실행결과를 볼 수는 있다.
	// fmt.Println("fin")
	// say("End")
}
