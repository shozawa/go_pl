package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
	count := make(chan struct{})
	finish := make(chan struct{})

	go func() {
		for {
			select {
			case <-finish:
				return
			default:
				ping <- struct{}{}
				count <- struct{}{}
				<-pong
			}
		}
	}()

	go func() {
		for {
			select {
			case <-finish:
				return
			default:
				<-ping
				pong <- struct{}{}
			}
		}
	}()

	var n int

	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			close(finish)
			fmt.Println(n)
			return
		case <-count:
			n++
		}
	}
}
