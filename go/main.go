package main

import (
	"fmt"
	"time"
)

type bus struct {
	a chan int
}

func (b *bus) run() {
	go func() {
		for {
			select {
			case n2 := <-b.a:
				fmt.Printf("%d start.\n", n2)
				fmt.Printf("%d done.\n", n2)
			}
		}
	}()
}

func main() {
	run2()
}

func run1() {
	b := bus{
		a: make(chan int),
	}
	b.run()
	for i := 0; i < 1000; i++ {
		go func() {
			b.a <- i
		}()
	}
	time.Sleep(time.Minute)
}

func run2() {
	b := bus{
		a: make(chan int),
	}
	b.run()
	for i := 0; i < 1000; i++ {
		go func(n int) {
			b.a <- n
		}(i)
	}
	time.Sleep(time.Minute)
}
