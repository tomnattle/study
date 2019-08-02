package main

import (
	"context"
	"fmt"

	"time"
)

func doJob(ctx context.Context, c <-chan int) (cancel context.CancelFunc) {
	fmt.Println("hi do job")
	defer func() {
		fmt.Println("exit do job")
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return
		case <-c:
			fmt.Println("recieve a job")
			//default:
			//	fmt.Println("===")
		}
	}
	//return cancel
}

func main() {
	cc := make(chan int, 0)
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)
	go doJob(ctxCancel, cc)
	go func() {
		for {
			<-time.Tick(1 * time.Second)
			cc <- 1
		}
	}()

	time.AfterFunc(10*time.Second, cancel)
	fmt.Println("nice")
	c := make(chan int, 0)
	c1 := make(chan int, 0)
	go job(c, c1)

	go func() {
		time.Sleep(12 * time.Second)
		c <- 1
	}()
	<-c1
	deferCall()
}

func job(c chan int, c1 chan int) {
	defer func() {
		fmt.Println("exit")
	}()
	for {
		select {
		case <-c:
			c1 <- 1
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("...")
		}
	}
}

func deferCall() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	type student struct {
		name string
		age  int
	}
	stus := []student{
		{name: "a", age: 1},
		{name: "b", age: 2},
	}
	for _, x := range stus {
		x.age += 1
	}
	fmt.Printf(":: %#+v\n", stus)
	return
	//panic("触发异常")
}
