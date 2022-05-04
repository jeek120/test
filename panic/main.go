package main

import (
	"fmt"
	"time"
)

type actor struct {
	e chan int
}

func main() {
	a := actor{e: make(chan int, 10)}
	a.run3()
	go func() {
		for {
			a.e <- 1
			fmt.Println("发送了1个")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Hour)
}

func (a *actor) run1() {
	go func() {
		for {
			select {
			case <-a.e:
				fmt.Printf("收到一个消息")
				panic(1)
			}
		}
	}()
}

func (a *actor) run3() {
	go func() {
		for {
			func() {
				defer func() {
					err := recover()
					if err != nil {
						fmt.Printf("错误恢复: %v", err)
					}
				}()
				fmt.Println("开始了run1")
				select {
				case <-a.e:
					fmt.Printf("收到一个消息")
					panic(1)
				}
			}()
		}
	}()
}
